package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	go consumer(ctx, "1")
	// go consumer(ctx, "2")
	// go consumer(ctx, "3")

	time.Sleep(time.Second * 100)
}

func consumer(ctx context.Context, consumerName string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	for {
		// Use BLMOVE to atomically move an item from "queue" to "processing"
		result, err := client.BLMove(ctx, "queue", "processing", "RIGHT", "LEFT", 0).Result()
		if err != nil {
			fmt.Printf("[%s] error moving from queue to processing: %v\n", consumerName, err)
			continue
		}

		fmt.Printf("[%s] processing: %v\n", consumerName, result)

		// Wait a random amount of time before marking the item as done
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

		// Move item from "processing" to "done"
		_, err = client.LRem(ctx, "processing", 1, result).Result()
		if err != nil {
			fmt.Printf("[%s] error removing from processing: %v\n", consumerName, err)
			continue
		}

		_, err = client.RPush(ctx, "done", result).Result()
		if err != nil {
			fmt.Printf("[%s] error pushing to done: %v\n", consumerName, err)
			continue
		}

		fmt.Printf("[%s] moved to done: %v\n", consumerName, result)
	}
}

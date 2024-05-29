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
		result, err := client.BRPop(ctx, 0, "queue").Result()
		if err != nil {
			fmt.Printf("[%s] error popping from queue: %w\n", consumerName, err)
			continue
		}

		fmt.Printf("[%s] received: %v\n", consumerName, result)

		// Wait a random amount of time before popping the next item
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

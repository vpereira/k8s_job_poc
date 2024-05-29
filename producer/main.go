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

	go producer(ctx)

	// Prevent the program to exit
	time.Sleep(time.Second * 100)
}

// producer pushes a new random item to the queue every second
func producer(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			// Push a random integer in the queue
			v := rand.Int()

			_, err := client.LPush(ctx, "queue", v).Result()
			if err != nil {
				fmt.Printf("error pushing to queue: %w", err)
			} else {
				fmt.Printf("pushed %d to queue\n", v)
			}
		}
	}
}

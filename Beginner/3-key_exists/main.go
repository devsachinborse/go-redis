package main 

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)


func main() {
    // Create a new Redis client
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis server address
        Password: "",               // No password set
        DB:       0,                // Use default DB
    })

    ctx := context.Background()
    key := "mykey"

    // Check if the key exists
    exists, err := rdb.Exists(ctx, key).Result()
    if err != nil {
        panic(err)
    }

    if exists > 0 {
        fmt.Printf("Key '%s' exists in Redis.\n", key)
    } else {
        fmt.Printf("Key '%s' does not exist in Redis.\n", key)
    }
}
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

func main() {
    // Create a new Redis client
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", 
        Password: "",               
        DB:       0,              //default DB
    })

    ctx := context.Background()
    key := "foo"

    // Set the key with some value for demonstration purposes
    err := rdb.Set(ctx, key, "bar", 0).Err()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Key '%s' set successfully.\n", key)

    // Set an expiration time of 10 seconds for the key
    err = rdb.Expire(ctx, key, 10*time.Second).Err()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Expiration time set for key '%s'.\n", key)

    // Wait for 11 seconds to check if the key expires
    fmt.Println("Waiting for 10 seconds...")
    time.Sleep(10 * time.Second)

    // Check if the key still exists
    exists, err := rdb.Exists(ctx, key).Result()
    if err != nil {
        panic(err)
    }

    if exists > 0 {
        fmt.Printf("Key '%s' still exists in Redis.\n", key)
    } else {
        fmt.Printf("Key '%s' has expired and no longer exists in Redis.\n", key)
    }
}

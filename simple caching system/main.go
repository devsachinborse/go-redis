package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/go-redis/redis/v8"
)

//user struct
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

//fetch a use data(simulate)
func fetchUserFromDB(userID int) *User {
	
	//simulated record
	users := map[int]*User {
		1:{ID:1, Name:"John", Age:30},
		2:{ID:2, Name:"Jane", Age:35},
	}

	return users[userID]

}


func main() {
	//create redis client
	rdb := redis.NewClient(&redis.Options {
		Addr:"localhost:6379",
		Password:"",
		DB:0, //default db
	})

	ctx := context.Background()
	userID := 2
	 //useID example
	

	//try to get user from cache
	cachedUserData, err := rdb.Get(ctx, fmt.Sprintf("user:%d", userID)).Result()
    if err == redis.Nil {
        fmt.Println("Cache miss. Fetching from database...")
        // Cache miss: fetch user data from the "database"
        user := fetchUserFromDB(userID)
        if user != nil {
            // Serialize user data to JSON
            userData, _ := json.Marshal(user)
            // Store user data in Redis with an expiration time of 10 seconds
            err := rdb.Set(ctx, fmt.Sprintf("user:%d", userID), userData, 10*time.Second).Err()
            if err != nil {
                panic(err)
            }
            fmt.Printf("User data cached: %+v\n", user)
        } else {
            fmt.Println("User not found in database.")
        }
    } else if err != nil {
        panic(err)
    } else {
        // Cache hit: deserialize the cached user data
        var cachedUser User
        json.Unmarshal([]byte(cachedUserData), &cachedUser)
        fmt.Printf("Cache hit: %+v\n", cachedUser)
    }
}


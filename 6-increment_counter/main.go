package main 

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	//create a new redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	ctx := context.Background()
	key := "foo"


	//increment the counter
	newVal , err := rdb.Incr(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The new value of '%s' is : %d\n", key, newVal)
}
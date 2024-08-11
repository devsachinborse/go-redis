package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main(){

	//create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",   //redis server port
		Password:"",
		DB:0, //use default DB
	})

	ctx := context.Background()


	//set the value

	err := rdb.Set(ctx, "hello", "world",0).Err()
	if err != nil {
		log.Fatal(err)
	}

	//get the value
	val ,err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello", val)
}
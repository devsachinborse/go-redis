package main 

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main () {
	//create a new redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	ctx := context.Background()
	key := "mystring"


	//set key
	err := rdb.Set(ctx, key, "foo" ,0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The key '%s' set successfully.....\n" , key)


	//get the string value
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The value of %s is %s\n", key, val)


	//append to string
	newVal, err := rdb.Append(ctx, key , "bar").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("appended to %s. New length: %d\n", key, newVal)


	//get updated key
	updatedkey , err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The updated value of %s is %s\n", key, updatedkey)

	//get the length of the string
	strLen, err := rdb.StrLen(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The length of %s is %d\n", key,strLen)
}
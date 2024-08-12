package main 


import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main(){
	//create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr : "localhost:6379", //redis server
		Password:"",
		DB:0,
	}) 

	ctx := context.Background()
	key := "foo"


	//set the key
	err := rdb.Set(ctx, key, "Bar",0).Err()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("key '%s' set successfully.\n", key)

	time.Sleep(2* time.Second) 

	//delete key
	delete , err := rdb.Del(ctx,key).Result()
	if err != nil {
		panic(err)
	}

	if delete > 0 {
		fmt.Printf("Key %s deleted successflly.\n",key)
	}else {
		fmt.Printf("key %s does not exists.", key)
	}
}


package main 

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	//create new redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:"localhost:6379", //redis server
		Password:"",
		DB:0, //default
	})

	ctx := context.Background()
	key := "mylist"

	//push elemets to the list (list-push)
	rdb.LPush(ctx, key, "one")
	rdb.LPush(ctx, key, "two")
	rdb.LPush(ctx, key, "three")
	rdb.LPush(ctx, key, "four")
	rdb.LPush(ctx, key, "five")

	fmt.Printf("Elements pushed to the list : %s \n", key)




	//get the elements from the list
	list, err := rdb.LRange(ctx, key, 0, -1).Result() //Here, 0 and -1 are used to retrieve all elements from the list
	if err != nil {
		panic(err)
	}
	fmt.Printf("The elements in the  '%s' are: %v\n", key, list)

	//get the length of the list
	len , err := rdb.LLen(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The elemets in the  %s are %d\n", key, len)

	//pop an element from the list(left-pop)
	popped , err := rdb.LPop(ctx,key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Element '%s' popped from the list '%s'.\n", popped, key)


	// Get the updated list
    updatedList, err := rdb.LRange(ctx, key, 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("The updated elements in the  '%s' are: %v\n", key, updatedList)
}
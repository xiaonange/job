package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main()  {
	client := redis.NewClient(&redis.Options{
		Addr:     "121.37.247.194:6379",
		Password: "Shijian@123456", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	client.HGetAll()


}

func TestString(client redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
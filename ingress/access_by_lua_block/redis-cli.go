package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("golang连接redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}

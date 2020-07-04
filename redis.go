package main

import (
	"fmt"
	redis "github.com/go-redis/redis/v7"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	get(redisCli)
}

func get(cli *redis.Client) {
	resp, err := cli.Get("testkey").Result()
	if err != nil {
		fmt.Printf("get err:%+v\n", err)
		return
	}
	fmt.Printf("resp:%+v, err:%+v\n", resp, err)
}

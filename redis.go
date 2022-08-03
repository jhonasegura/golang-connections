package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func() {

	// new redis client
	client := redis.NewClient(&redis.Options{

		Addr: "127.0.0.1:6379",

		Password: "",

		DB: 10,
	})

	// test connection
	pong, err := client.Ping().Result()

	if err != nil {

		log.Fatal(err)

	}

	// return pong if server is online
	fmt.Println(pong)
}

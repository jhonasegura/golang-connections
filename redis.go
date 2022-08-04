package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/genproto/googleapis/cloud/redis/v1"
)

type Employee struct {
	Name    string `json:"name`
	Address string `json:"address`
}

func main() {

	// new redis client
	client := redis.NewClient(&redis.Options{

		Addr:     "localhost:6000",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Employee{Name: "Gordon", Address: "San Jose"})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("add123", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("add123").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

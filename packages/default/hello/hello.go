package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	fmt.Println("parsing redis url")
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	fmt.Println("creating redis client")
	client := redis.NewClient(opts)

	fmt.Println("pinging")
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("ponging")

	msg["body"] = fmt.Sprintf("%s. Ping result: %s", msg["body"], pong)

	return msg
}

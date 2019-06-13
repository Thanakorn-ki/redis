package main

import (
	"fmt"

	"github.com/salapao2136/redis/app"
)

func main() {
	redisdb := app.NewRedisClient()
	// Publish a message.
	err := redisdb.Publish("mychannel1", "mychannel1").Err()
	if err != nil {
		panic(err)
	}

	err = redisdb.Publish("mychannel", "mychannel").Err()
	if err != nil {
		panic(err)
	}

	fmt.Print("Publish Start")
}

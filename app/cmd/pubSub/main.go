package main

import (
	"fmt"
	"time"

	"github.com/salapao2136/redis/app"
)

func main() {
	redisdb := app.NewRedisClient()
	pubsub := redisdb.PSubscribe("*")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

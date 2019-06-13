package app

import (
	"log"
	"os"

	goRedis "github.com/go-redis/redis"
)

// NewRedisClient connetions redis simple
func NewRedisClient() *goRedis.Client {
	redisdb := goRedis.NewClient(&goRedis.Options{
		Addr:     os.Getenv("REDIS_CLUSTER_ADDR"), // use default Addr
		Password: "",                              // no password set
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		log.Print("Unable connect redis client")
		log.Panic(err)
	}

	return redisdb
}

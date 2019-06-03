package redis

import (
	"github.com/go-redis/redis"
	"fmt"
)

var RedisClient *redis.Client

func init() {
	fmt.Println("redis ...")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
}

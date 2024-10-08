package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	RDB = InitRedis()
	Ctx = context.Background()
)

func InitRedis() *redis.Client {
	log.Printf("Init redis")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Cannot connect to redis: %s\n", err)
	}
	return rdb
}

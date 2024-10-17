package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg *config.RedisConfig) (*redis.Client, error) {
	addrString := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addrString,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Cannot connect to redis: %s\n", err)
		return nil, err
	}
	return rdb, nil
}

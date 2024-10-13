package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	"github.com/go-redis/redis/v8"
)

var (
	cfg = config.GetConfig().Redis
	RDB = InitRedis(&cfg)
	Ctx = context.Background()
)

func InitRedis(cfg *config.RedisConfig) *redis.Client {
	log.Printf("Init redis")
	addrString := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addrString,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Cannot connect to redis: %s\n", err)
	}
	return rdb
}

package redis

import (
	"context"
	"strings"
	"testing"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
)

func TestInitRedis(t *testing.T) {

	Config := &config.RedisConfig{
		Host: Host,
		Port: Port,
	}

	rdb, err := InitRedis(Config)
	if err != nil {
		t.Errorf("InitRedis failed: %s", err)
	}

	if rdb == nil {
		t.Errorf("Redis connection is not established")
	}

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		t.Errorf("Redis ping failed: %s", err)
	}
}

func TestInitRedisError(t *testing.T) {

	fakeConfig := &config.RedisConfig{
		Host: FakeHost,
		Port: Port,
	}

	_, err := InitRedis(fakeConfig)
	if err == nil {
		t.Errorf("InitRedis should return an error")
	}

}

func TestInitRedisErrorPort(t *testing.T) {
	fakeConfig := &config.RedisConfig{
		Host: Host,
		Port: FakePort,
	}

	_, err := InitRedis(fakeConfig)
	if err == nil {
		t.Errorf("InitRedis should return an error")
	}

	if !strings.Contains(err.Error(), "dial tcp") {
		t.Errorf("InitRedis should return a connection error, but got: %v", err)
	}
}

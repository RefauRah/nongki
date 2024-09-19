package config

import (
	"context"
	"nongki/pkg/log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	logger := log.GetLogger()
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:   0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(err)
	}
}

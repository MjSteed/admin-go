package common

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Cache *redis.Client

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		LOG.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}

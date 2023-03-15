package bootstrap

import (
	"context"
	"fmt"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/redis/go-redis/v9"
)

func InitializeRedis() *redis.Client {
	config := common.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("redis connect ping failed: %s ", err))
	}
	return client
}

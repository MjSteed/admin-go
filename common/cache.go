package common

import (
	"context"
	"encoding/json"
	"time"

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

// 从缓存获取数据转化为需要的类型
func CacheGet(key string, data interface{}) error {
	str, err := Cache.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	json.Unmarshal([]byte(str), &data)
	return nil
}

func CacheSet(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = Cache.Set(context.Background(), key, data, time.Second*60).Err()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}

func CacheDel(key ...string) error {
	c, err := Cache.Del(context.Background(), key...).Result()
	LOG.Debug("已删除key数量", zap.Int64("count", c), zap.Any("key", key))
	if err == redis.Nil {
		LOG.Debug("要删除的key不存在", zap.Any("key", key))
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const DefaultExpireTime = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
	client *redis.Client
}

func InitRedis() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("Redis loaded successfully")
	return &RedisClient{
		client: client,
	}, nil
}

func (r *RedisClient) Set(key string, value any, rest ...any) error {
	expireTime := DefaultExpireTime
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			expireTime = v
		}
	}
	return r.client.Set(context.Background(), key, value, expireTime).Err()
}

func (r *RedisClient) Get(key string) (any, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisClient) Delete(key ...string) error {
	return r.client.Del(context.Background(), key...).Err()
}

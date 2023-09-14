package utils

import (
	config "auth-service/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// NewRedisClient creates a new redis client
func NewRedisClient(config config.ConfigType) *redis.Client{
	if config.Redis.Address == "" {
		return nil
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
		PoolSize: config.Redis.PoolSize,
		MinIdleConns: config.Redis.MinIdleConns,
	})
	return RedisClient
}
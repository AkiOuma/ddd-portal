package redisdb

import (
	"janus-portal/config/persistence"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(config *persistence.RedisDBConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr(),
		Password: config.Password(),
		DB:       config.DB(),
	})
}

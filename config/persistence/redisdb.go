package persistence

import (
	"os"
	"strconv"
)

type RedisDBConfig struct {
	host     string
	port     string
	password string
	db       string
}

func NewRedisDBConfig() *RedisDBConfig {
	return &RedisDBConfig{
		host:     os.Getenv("redis_host"),
		port:     os.Getenv("redis_port"),
		password: os.Getenv("redis_password"),
		db:       os.Getenv("redis_db"),
	}
}

func (config *RedisDBConfig) Addr() string {
	return config.host + ":" + config.port
}

func (config *RedisDBConfig) Password() string {
	return config.password
}

func (config *RedisDBConfig) DB() int {
	db, err := strconv.Atoi(config.db)
	if err != nil {
		db = 0
	}
	return db
}

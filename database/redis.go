package database

import (
	"github.com/redis/go-redis/v9"
	"smash/config"
)

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
	})
}

package redis

import "github.com/go-redis/redis/v8"

var RedisClient *redis.Client

func NewRedisClient(host string, password string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
}

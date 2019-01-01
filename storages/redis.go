package storages

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
}

var once sync.Once
var redisClient *RedisClient

func GetRedisClient() *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		redisClient = &RedisClient{client}
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}

	return redisClient
}

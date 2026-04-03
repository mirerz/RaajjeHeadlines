package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Redis connection established.")
	}
}

func SetCache(key string, value string, ttl time.Duration) error {
	return RDB.Set(ctx, key, value, ttl).Err()
}

func GetCache(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func ClearCache(key string) error {
	return RDB.Del(ctx, key).Err()
}

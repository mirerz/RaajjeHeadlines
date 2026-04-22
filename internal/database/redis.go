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
		RDB = nil // Ensure RDB is nil if connection failed
	} else {
		log.Println("Redis connection established.")
	}
}

func SetCache(key string, value string, ttl time.Duration) error {
	if RDB == nil {
		return nil
	}
	return RDB.Set(ctx, key, value, ttl).Err()
}

func GetCache(key string) (string, error) {
	if RDB == nil {
		return "", redis.Nil
	}
	return RDB.Get(ctx, key).Result()
}

func ClearCache(key string) error {
	if RDB == nil {
		return nil
	}
	return RDB.Del(ctx, key).Err()
}

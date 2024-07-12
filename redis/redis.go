package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rbd *redis.Client

func Init() {
	rbd = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rbd.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func Get(key string) string {
	val, err := rbd.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Failed to get value from Redis: %v", err)
	}
	return val
}

func Set(key string, value string) {
	err := rbd.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set value to Redis: %v", err)
	}
}

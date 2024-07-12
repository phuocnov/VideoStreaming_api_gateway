package redis

import (
	"APIGateway/pkg/dto"
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func GetTodos(key string, todos *[]dto.Todo) error {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), todos); err != nil {
		return err
	}

	return nil
}

func SetTodos(key string, todos []dto.Todo) error {
	val, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	if err := rdb.Set(ctx, key, val, 0).Err(); err != nil {
		return err
	}

	return nil
}

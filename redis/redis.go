package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Adjust if needed
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")
}

func Set(key string, value string, expiration time.Duration) error {
	return RDB.Set(Ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return RDB.Get(Ctx, key).Result()
}

func Del(key string) error {
	return RDB.Del(Ctx, key).Err()
}

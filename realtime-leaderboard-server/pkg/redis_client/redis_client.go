package redis_client

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Test the connection
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	if rdb == nil {
		InitRedis()
	}
	return rdb
}

func GetCtx() context.Context {
	return ctx
}

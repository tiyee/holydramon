package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(addr, password string) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	var ctx = context.Background()
	return RedisClient.Ping(ctx).Err()

}

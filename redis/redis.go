package redis

import (
	"context"
	"goparkin_service/config"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: config.AppConfig.RedisAddr,
    })
    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatal(err)
    }
}
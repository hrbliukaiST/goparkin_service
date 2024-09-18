go
package redis

import (
    "github.com/go-redis/redis/v8"
    "log"
    "myproject/config"
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
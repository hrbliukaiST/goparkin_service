go
package config

import (
    "log"
    "os"
)

type Config struct {
    MySQLDSN   string
    MongoDBURI string
    RedisAddr  string
    MQTTBroker string
    EmailUser  string
    EmailPass  string
}

var AppConfig Config

func LoadConfig() {
    AppConfig = Config{
        MySQLDSN:   os.Getenv("MYSQL_DSN"),
        MongoDBURI: os.Getenv("MONGODB_URI"),
        RedisAddr:  os.Getenv("REDIS_ADDR"),
        MQTTBroker: os.Getenv("MQTT_BROKER"),
        EmailUser:  os.Getenv("EMAIL_USER"),
        EmailPass:  os.Getenv("EMAIL_PASS"),
    }

    if AppConfig.MySQLDSN == "" || AppConfig.MongoDBURI == "" {
        log.Fatal("Missing required environment variables")
    }
}
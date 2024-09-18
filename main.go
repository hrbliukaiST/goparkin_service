go
package main

import (
    "log"
    "net/http"
    "myproject/config"
    "myproject/db"
    "myproject/mqtt"
    "myproject/redis"
    "myproject/routes"
)

func main() {
    config.LoadConfig()
    db.InitMySQL()
    db.InitMongoDB()
    mqtt.InitMQTT()
    redis.InitRedis()

    router := routes.SetupRoutes()

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
package main

import (
	"goparkin_service/config"
	"goparkin_service/db"
	"goparkin_service/mqtt"
	"goparkin_service/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "goparkin_service/docs"
)

// Change to your project path

// @title Goparkin API
// @version 1.0
// @description This is a goparkin back end API.
// @host localhost:7080
// @BasePath /

func main() {

	log.Println("Server started ")
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	 // Check for required variables
	 requiredVars := []string{"MYSQL_DSN", "MONGODB_URI"} // Replace with your actual variable names
	 for _, v := range requiredVars {
		 if os.Getenv(v) == "" {
			 log.Fatalf("Missing required environment variable: %s", v)
		 }
	 }	 		
	config.LoadConfig()
	log.Println("File loaded")
    db.InitMySQL()
    // db.InitMongoDB()
    mqtt.InitMQTT()
    // redis.InitRedis()


    router := routes.SetupRoutes()

    log.Println("Server started at :7080")
    if err := http.ListenAndServe(":7080", router); err != nil {
        log.Fatal(err)
    }
}



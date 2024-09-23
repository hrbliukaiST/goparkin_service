package routes

import (
	_ "goparkin_service/docs" // Import the generated docs
	"goparkin_service/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	// http-swagger middleware
)

func SetupRoutes() *mux.Router {
   
    // router := gin.Default()

    // // Swagger endpoint
    // router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // // Define your routes
    // router.POST("/users", gin.WrapF(handlers.CreateUser))
    // router.GET("/users/:id", func(c *gin.Context) {
    //     handlers.GetUserByID(c.Writer, c.Request)
    // })
    router := mux.NewRouter()

    // Add CREATE user route
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

    // // Add GET user route
    // router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET") // {{ edit_1 }}

    // Swagger UI
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    return router
}
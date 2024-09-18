package routes

import (
	_ "goparkin_service/docs" // Import the generated docs
	"goparkin_service/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

    // Swagger UI
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    return router
}
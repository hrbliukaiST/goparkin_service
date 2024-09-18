go
package routes

import (
    "github.com/gorilla/mux"
    "myproject/handlers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    return router
}
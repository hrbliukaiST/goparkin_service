package routes

import (
	_ "goparkin_service/docs" // Import the generated docs
	"goparkin_service/handlers"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	// http-swagger middleware
)

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte(`{"error": "Method Not Allowed"}`))
}

func SetupRoutes() *mux.Router {
   
  
    router := mux.NewRouter()

    // Swagger UI
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    // Add CREATE user route
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

    // // Add GET user route
    router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")

    // Add Update user route
    router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

    // Handle 405 Method Not Allowed
    router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)


    return router
}
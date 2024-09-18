package handlers

import (
	"encoding/json"
	"goparkin_service/models"
	"net/http"
	// Swagger packages
	// Import docs
)

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User info"
// @Success 200 {object} models.User
// @Failure 400 {object} string
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Here you would typically save user in MySQL or MongoDB
    // For this example, we will return the user object
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
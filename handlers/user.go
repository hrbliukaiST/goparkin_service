package handlers

import (
	"encoding/json"
	"goparkin_service/db"
	"goparkin_service/models"
	"goparkin_service/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// Swagger packages
	// Import docs
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} string
// @Router /users [post]

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	log.Println("userIDStr loaded",userIDStr)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := db.CheckAndCreateTable("users"); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	if err := service.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Here you would typically save user in MySQL or MongoDB
	// For this example, we will return the user object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
/*
// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieve a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} string
// @Router /users/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract userID from the URL
	vars := mux.Vars(r)
    userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}*/
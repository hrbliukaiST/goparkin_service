package handlers

import (
	"encoding/json"
	"goparkin_service/db"
	"goparkin_service/models"
	"goparkin_service/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// Swagger packages
	// Import docs
)

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
	log.Println("Request :",r)	
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
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update user information by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User details"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} string
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Extract userID from the URL
	log.Println("Request :",r)	
    vars := mux.Vars(r)
    userIDStr := vars["id"]
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
	log.Println("user :",userID, user)	

    updatedUser, err := service.UpdateUser(userID, user)
    if err != nil {
        if err == service.ErrUserNotFound {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to update user", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedUser)
}

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

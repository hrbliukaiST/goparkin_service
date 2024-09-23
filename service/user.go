package service

import (
	"goparkin_service/db"
	"goparkin_service/models"
)

// GetUserByID retrieves a user by their unique ID from the database
func GetUserByID(userID int) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	// Use MySQLDB here
	err := db.MySQLDB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err // Return nil and the error if user is not found
	}
	return &user, nil // Return the user if found
}

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.MySQLDB.Exec(query, user.Name, user.Email) // Adjust fields as necessary
	return err // Return the error if any
}
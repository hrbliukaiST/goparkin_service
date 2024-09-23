package service

import (
	"errors"
	"goparkin_service/db"
	"goparkin_service/models"
	"log"
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

var ErrUserNotFound = errors.New("user not found")

// UpdateUser updates an existing user in the database
func UpdateUser(userID int, user models.User) (models.User, error) {

    log.Println("user :",userID, user)	
    query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
    result, err := db.MySQLDB.Exec(query, user.Name, user.Email, userID)
    if err != nil {
        return models.User{}, err
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return models.User{}, err
    }
    if rowsAffected == 0 {
        return models.User{}, ErrUserNotFound
    }
    user.ID = userID
    return user, nil


}


// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.MySQLDB.Exec(query, user.Name, user.Email) // Adjust fields as necessary
	return err // Return the error if any
}
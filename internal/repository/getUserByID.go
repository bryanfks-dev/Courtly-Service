package repository

import (
	"main/data/models"
	"main/internal/providers/database"
)

// GetUserByID is a repository function that returns the user with the given ID.
//
// id: The ID of the user.
//
// Returns the user with the given ID and an error if any.
func GetUserByID(id uint) (models.User, error) {
	// Create a new user instance
	var user models.User

	// Find the user with the given ID
	err := database.Conn.Where("id = ?", id).First(&user).Error

	// Return an error if any
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

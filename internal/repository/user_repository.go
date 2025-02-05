package repository

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"
)

// UserRepository is a struct that defines the user repository.
type UserRepository struct{}

// NewUserRepository is a factory function that returns a new instance of the user repository.
//
// Returns a new instance of the user repository.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create is a function that creates a new user.
//
// user: The user object.
//
// Returns an error if any.
func (*UserRepository) Create(user *models.User) error {
	// Create the user
	err := mysql.Conn.Create(user).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to create user: " + err.Error())

		return err
	}

	return nil
}

// GetUsingID is a function that returns a user by ID.
//
// userID: The user ID.
//
// Returns the user object and an error if any.
func (*UserRepository) GetUsingID(userID uint) (*models.User, error) {
	// Create a new user object
	var user models.User

	// Get the user by ID
	err := mysql.Conn.First(&user, "id = ?", userID).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to get user using id: " + err.Error())

		return nil, err
	}

	return &user, err
}

// GetUsingUsername is a function that returns a user by username.
//
// username: The username.
//
// Returns the user object and an error if any.
func (*UserRepository) GetUsingUsername(username string) (*models.User, error) {
	// Create a new user object
	var user models.User

	// Get the user by username
	err := mysql.Conn.First(&user, "username = ?", username).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to get user using username: " + err.Error())

		return nil, err
	}

	return &user, err
}

// GetUsinPhoneNumber is a function that returns a user by phone number.
func (*UserRepository) GetUsingPhoneNumber(phoneNumber string) (*models.User, error) {
	// Create a new user object
	var user models.User

	// Get the user by phone number
	err := mysql.Conn.First(&user, "phone_number = ?", phoneNumber).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to get user using phone number: " + err.Error())

		return nil, err
	}

	return &user, err
}

// IsUsernameTaken is a function that checks if a username is taken.
//
// username: The username.
//
// Returns a boolean indicates username is taken and an error if any.
func (*UserRepository) IsUsernameTaken(username string) (bool, error) {
	// Create a counter variable
	var count int64

	// Check if the username is taken
	err := mysql.Conn.Model(&models.User{}).Where("username = ?", username).Limit(1).Count(&count).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to check if username is taken: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

// IsPhoneNumberTaken is a function that checks if a phone number is taken.
//
// phoneNumber: The phone number.
//
// Returns a boolean indicates phone number is taken and an error if any.
func (*UserRepository) IsPhoneNumberTaken(phoneNumber string) (bool, error) {
	// Create a counter variable
	var count int64

	// Check if the phone number is taken
	err := mysql.Conn.Model(&models.User{}).Where("phone_number = ?", phoneNumber).Limit(1).Count(&count).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to check if phone number is taken: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

// UpdatePassword is a function that updates a user's password.
//
// userID: The user ID.
// hashedNewPassword: The hashed new password.
//
// Returns an error if any.
func (*UserRepository) UpdatePassword(userID uint, hashedNewPassword string) error {
	// Update the user's password
	err := mysql.Conn.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedNewPassword).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to update user's password: " + err.Error())

		return err
	}

	return nil
}

// UpdateUsername is a function that updates a user's username.
//
// userID: The user ID.
// newUsername: The new username.
//
// Returns an error if any.
func (*UserRepository) UpdateUsername(userID uint, newUsername string) error {
	// Update the user's username
	err := mysql.Conn.Model(&models.User{}).Where("id = ?", userID).Update("username", newUsername).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to update user's username: " + err.Error())

		return err
	}

	return nil
}

// UpdateProfilePicture is a function that updates a user's profile picture.
//
// userID: The user ID.
// newFileName: The new file name.
//
// Returns an error if any.
func (*UserRepository) UpdateProfilePicture(userID uint, newFileName string) error {
	// Update the user's username
	err := mysql.Conn.Model(&models.User{}).Where("id = ?", userID).Update("profile_picture", newFileName).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to update user's profile picture: " + err.Error())

		return err
	}

	return nil
}

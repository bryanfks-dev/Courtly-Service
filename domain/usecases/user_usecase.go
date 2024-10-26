package usecases

import (
	"log"
	"main/data/models"
	"main/internal/repository"
)

// UserUseCase is a struct that defines the use case for the user entity.
type UserUseCase struct {
	userRepository *repository.UserRepository
}

// NewUserUseCase is a factory function that returns a new instance of the UserUseCase struct.
//
// u: The user repository.
//
// Returns a new instance of the UserUseCase.
func NewUserUseCase(u *repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: u}
}

// GetUserByID is a method that returns a user entity by its ID.
//
// id: The ID of the user.
//
// Returns the user entity and an error if any.
func (u *UserUseCase) GetUserByID(id uint) (*models.User, error) {
	// Get the user entity by its ID
	user, err := u.userRepository.GetUsingID(id)

	// Check if there is an error
	if err != nil {
		log.Println("Failed retrieve user by ID: ", err)

		return nil, err
	}

	return user, nil
}

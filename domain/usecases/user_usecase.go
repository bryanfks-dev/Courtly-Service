package usecases

import (
	"log"
	"main/data/models"
	"main/domain/entities"
	"main/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// UserUseCase is a struct that defines the use case for the user entity.
type UserUseCase struct {
	AuthUseCase    *AuthUseCase
	UserRepository *repository.UserRepository
}

// NewUserUseCase is a factory function that returns a new instance of the UserUseCase struct.
//
// a: The auth use case.
// u: The user repository.
//
// Returns a new instance of the UserUseCase.
func NewUserUseCase(a *AuthUseCase, u *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		AuthUseCase:    a,
		UserRepository: u,
	}
}

// GetUserUsingID is a method that returns the user using the given ID.
//
// userID: The ID of the user.
//
// Returns the user and an error if any.
func (u *UserUseCase) GetUserUsingID(userID uint) (*models.User, *entities.ProcessError) {
	// Get the user from the database
	user, err := u.UserRepository.GetUsingID(userID)

	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			Message:     "User not found",
			ClientError: true,
		}
	}

	// Return an error if any
	if err != nil {
		log.Println("Failed to get current user: ", err)

		return nil, &entities.ProcessError{
			Message:     "Failed to get current user",
			ClientError: false,
		}
	}

	return user, nil
}

// GetCurrentUser is a method that returns the current user.
func (u *UserUseCase) GetCurrentUser(token *jwt.Token) (*models.User, *entities.ProcessError) {
	// Get the token claims
	claims := u.AuthUseCase.DecodeToken(token)

	return u.GetUserUsingID(claims.Id)
}

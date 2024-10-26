package usecases

import (
	"fmt"
	"log"
	"main/core/constants"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"main/pkg/utils"
)

// ChangeUserUsernameUseCase is a struct that defines the use case for changing the user's username.
type ChangeUserUsernameUseCase struct {
	userRepository *repository.UserRepository
}

// NewChangeUserUsernameUseCase is a factory function that returns a new instance of the ChangeUserUsernameUseCase struct.
//
// u: The user repository.
//
// Returns a new instance of the ChangeUserUsernameUseCase.
func NewChangeUserUsernameUseCase(u *repository.UserRepository) *ChangeUserUsernameUseCase {
	return &ChangeUserUsernameUseCase{userRepository: u}
}

// ValidateForm is a function that validates the change username form.
//
// form: The change username form.
//
// Returns an error if any.
func (c *ChangeUserUsernameUseCase) ValidateForm(form *dto.ChangeUsernameForm) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the username is empty
	if utils.IsBlank(form.NewUsername) {
		errs["username"] = append(errs["username"], "Username is required")
	}

	// Check if the username is too short
	if len(form.NewUsername) < constants.MINIMUM_USERNAME_LENGTH {
		errs["username"] = append(errs["username"], fmt.Sprintf("Username must be at least %d characters", constants.MINIMUM_USERNAME_LENGTH))
	}

	// Check if the errors map is not empty
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Process is a function that processes the change username use case.
//
// form: The change username form.
//
// Returns an error if any.
func (c *ChangeUserUsernameUseCase) Process(userID uint, form *dto.ChangeUsernameForm) (*models.User, *entities.ProcessError) {
	// Get the user by ID
	taken, err := c.userRepository.IsUsernameTaken(form.NewUsername)

	// Return an error if any
	if err != nil {
		log.Println("Failed to check if username is taken: ", err)

		return nil, &entities.ProcessError{
			Message:     "An error occurred while checking if the username is taken",
			ClientError: false,
		}
	}

	// Return an error if the username is taken
	if taken {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"The username is already taken"},
			},
			ClientError: true,
		}
	}

	// Update the username
	user, err := c.userRepository.UpdateUsername(userID, form.NewUsername)

	// Return an error if any
	if err != nil {
		log.Println("Failed to update username: ", err)

		return nil, &entities.ProcessError{
			Message:     "An error occurred while updating the username",
			ClientError: false,
		}
	}

	return user, nil
}

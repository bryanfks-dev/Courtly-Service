package usecases

import (
	"log"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"main/pkg/utils"
)

// LoginUseCase is a struct that defines the login use case.
type LoginUseCase struct {
	userRepository *repository.UserRepository
	authUseCase    *AuthUseCase
}

// NewLoginUseCase is a factory function that returns a new instance of the LoginUseCase.
//
// u: The user repository.
//
// Returns a new instance of the LoginUseCase.
func NewLoginUseCase(u *repository.UserRepository, a *AuthUseCase) *LoginUseCase {
	return &LoginUseCase{userRepository: u, authUseCase: a}
}

// ValidateForm is a function that validates the login form.
//
// form: The login form data.
//
// Returns a map of errors.
func (l LoginUseCase) ValidateForm(form *dto.LoginForm) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the username is blank
	if utils.IsBlank(form.Username) {
		errs["username"] = append(errs["username"], "Username is required")
	}

	// Check if the password is blank
	if utils.IsBlank(form.Password) {
		errs["password"] = append(errs["password"], "Password is required")
	}

	// Return the errors if any
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Process is a function that processes the login form.
//
// form: The login form data.
//
// Returns the user object and an error if any.
func (l LoginUseCase) Process(form *dto.LoginForm) (*models.User, *entities.ProcessError) {
	// Check if the username is taken
	user, err := l.userRepository.GetUsingUsername(form.Username)

	// Return an error if any
	if err != nil {
		log.Println("Error getting the user: ", err)

		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"An error occurred while getting the user"},
			},
			ClientError: true,
		}
	}

	// Return an error if the user does not exist
	if user == nil {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"Username does not exist"},
			},
			ClientError: true,
		}
	}

	// Check if the password is correct
	if !l.authUseCase.VerifyPassword(form.Password, user.Password) {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"password": []string{"Password is incorrect"},
			},
			ClientError: true,
		}
	}

	return user, nil
}

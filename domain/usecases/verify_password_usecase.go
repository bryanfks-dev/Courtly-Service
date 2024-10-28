package usecases

import (
	"log"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"main/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

// VerifyPasswordUseCase is a use case that provides the business logic for verifying a user's password.
type VerifyPasswordUseCase struct {
	AuthUseCase    *AuthUseCase
	UserRepository *repository.UserRepository
}

// VerifyPassword is a method of the VerifyPasswordUseCase that verifies a user's password.
//
// a: The AuthUseCase instance.
// u: The UserRepository instance.
//
// Retuns instance of the VerifyPasswordUseCase.
func NewVerifyPasswordUseCase(a *AuthUseCase, u *repository.UserRepository) *VerifyPasswordUseCase {
	return &VerifyPasswordUseCase{
		AuthUseCase:    a,
		UserRepository: u,
	}
}

// ValidateForm is a function that validates the verify password form.
//
// form: The verify password form data.
//
// Returns a map of errors.
func (*VerifyPasswordUseCase) ValidateForm(form *dto.VerifyPasswordForm) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

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

// Process is a function that processes the verify password form.
//
// form: The verify password form data.
// userID: The user ID.
//
// Returns the user object and an error if any.
func (v *VerifyPasswordUseCase) Process(form *dto.VerifyPasswordForm, token *jwt.Token) (*models.User, *entities.ProcessError) {
	// Decode the token
	claims := v.AuthUseCase.DecodeToken(token)

	// Get the user by the user ID
	user, err := v.UserRepository.GetUsingID(claims.Id)

	// Check if there is an error
	if err != nil {
		log.Println("Error getting user using ID: ", err)

		return nil, &entities.ProcessError{
			Message:     "An error occurred while getting the user",
			ClientError: false,
		}
	}

	// Check if the password is correct
	if !v.AuthUseCase.VerifyPassword(form.Password, user.Password) {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"password": []string{"Incorrect password"},
			},
			ClientError: true,
		}
	}

	return user, nil
}

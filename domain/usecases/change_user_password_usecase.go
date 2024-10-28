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

	"github.com/golang-jwt/jwt/v5"
)

// ChangeUserPasswordUseCase is a struct that defines the change password use case.
type ChangeUserPasswordUseCase struct {
	AuthUseCase    *AuthUseCase
	UserRepository *repository.UserRepository
}

// NewChangePasswordUseCase is a factory function that returns a new instance of the ChangePasswordUseCase.
//
// a: The auth use case.
// u: The user repository.
//
// Returns a new instance of the ChangePasswordUseCase.
func NewChangeUserPasswordUseCase(a *AuthUseCase, u *repository.UserRepository) *ChangeUserPasswordUseCase {
	return &ChangeUserPasswordUseCase{
		AuthUseCase:    a,
		UserRepository: u,
	}
}

// ValidateForm is a function that validates the change password form.
//
// form: The change password form.
//
// Returns a map of errors.
func (c *ChangeUserPasswordUseCase) ValidateForm(form *dto.ChangePasswordForm) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the old password is blank
	if utils.IsBlank(form.OldPassword) {
		errs["old_password"] = append(errs["old_password"], "Old password is required")
	}

	// Check if the new password is blank
	if utils.IsBlank(form.NewPassword) {
		errs["new_password"] = append(errs["new_password"], "New password is required")
	}

	if len(form.NewPassword) < constants.MINIMUM_PASSWORD_LENGTH {
		errs["new_password"] = append(errs["new_password"], fmt.Sprintf("Password must be at least %d characters long", constants.MINIMUM_PASSWORD_LENGTH))
	}

	// Check if the confirm password is blank
	if utils.IsBlank(form.ConfirmPassword) {
		errs["confirm_password"] = append(errs["confirm_password"], "Confirm password is required")
	}

	// Check if the new password and confirm password match
	if form.NewPassword != form.ConfirmPassword {
		errs["confirm_password"] = append(errs["confirm_password"], "Passwords do not match")
	}

	// Check if the errors map is not empty
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Process is a function that processes the change password use case.
//
// userID: The ID of the user.
// form: The change password form.
//
// Returns the user object and an error if any.
func (c *ChangeUserPasswordUseCase) Process(token *jwt.Token, form *dto.ChangePasswordForm) (*models.User, *entities.ProcessError) {
	// Get the user ID from the token
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the user by ID
	user, err := c.UserRepository.GetUsingID(claims.Id)

	// Check if there is an error
	if err != nil {
		log.Panicln("Error getting user: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while getting the user",
		}
	}

	// Check if the old password is correct
	if !c.AuthUseCase.VerifyPassword(form.OldPassword, user.Password) {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message: types.FormErrorResponseMsg{
				"old_password": []string{"Old password is incorrect"},
			},
		}
	}

	// Hash the new password
	hashedNewPassword, err := c.AuthUseCase.HashPassword(form.NewPassword)

	// Check if there is an error
	if err != nil {
		log.Println("Error hashing password: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while hashing the new password",
		}
	}

	// Update the user's password
	user, err = c.UserRepository.UpdatePassword(claims.Id, hashedNewPassword)

	// Check if there is an error
	if err != nil {
		log.Println("Error updating user's password: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while updating the user's password",
		}
	}

	return user, nil
}

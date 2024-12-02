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

// ValidateChangePasswordForm is a function that validates the change password form.
//
// form: The change password form dto.
//
// Returns a map of errors.
func (u *UserUseCase) ValidateChangePasswordForm(form *dto.ChangePasswordFormDTO) types.FormErrorResponseMsg {
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

// ProcessChangePassword is a function that processes the change password use case.
//
// token: The user token.
// form: The change password form dto.
//
// Returns an error if any.
func (u *UserUseCase) ProcessChangePassword(token *jwt.Token, form *dto.ChangePasswordFormDTO) *entities.ProcessError {
	// Get the user ID from the token
	claims := u.AuthUseCase.DecodeToken(token)

	// Get the user by ID
	user, err := u.UserRepository.GetUsingID(claims.Id)

	// Check if there is an error
	if err != nil {
		log.Panicln("Error getting user: ", err)

		return &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while getting the user",
		}
	}

	// Check if the old password is correct
	if !u.AuthUseCase.VerifyPassword(form.OldPassword, user.Password) {
		return &entities.ProcessError{
			ClientError: true,
			Message: types.FormErrorResponseMsg{
				"old_password": []string{"Old password is incorrect"},
			},
		}
	}

	// Hash the new password
	hashedNewPassword, err := u.AuthUseCase.HashPassword(form.NewPassword)

	// Check if there is an error
	if err != nil {
		log.Println("Error hashing password: ", err)

		return &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while hashing the new password",
		}
	}

	// Update the user's password
	err = u.UserRepository.UpdatePassword(claims.Id, hashedNewPassword)

	// Check if there is an error
	if err != nil {
		log.Println("Error updating user's password: ", err)

		return &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while updating the user's password",
		}
	}

	return nil
}

// ValidateChangeUsernameForm is a function that validates the change username form.
//
// form: The change username form dto.
//
// Returns an error if any.
func (u *UserUseCase) ValidateChangeUsernameForm(form *dto.ChangeUsernameFormDTO) types.FormErrorResponseMsg {
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

// ProcessUsername is a function that processes the change username use case.
//
// form: The change username form dto.
//
// Returns an error if any.
func (u *UserUseCase) ProcessChangeUsername(token *jwt.Token, form *dto.ChangeUsernameFormDTO) *entities.ProcessError {
	// Get the user by ID
	taken, err := u.UserRepository.IsUsernameTaken(form.NewUsername)

	// Return an error if any
	if err != nil {
		log.Println("Failed to check if username is taken: ", err)

		return &entities.ProcessError{
			Message:     "An error occurred while checking if the username is taken",
			ClientError: false,
		}
	}

	// Return an error if the username is taken
	if taken {
		return &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"The username is already taken"},
			},
			ClientError: true,
		}
	}

	// Get the user ID from the token
	claims := u.AuthUseCase.DecodeToken(token)

	// Update the username
	err = u.UserRepository.UpdateUsername(claims.Id, form.NewUsername)

	// Return an error if any
	if err != nil {
		log.Println("Failed to update username: ", err)

		return &entities.ProcessError{
			Message:     "An error occurred while updating the username",
			ClientError: false,
		}
	}

	return nil
}

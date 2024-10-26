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

// ChangeUserPasswordUseCase is a struct that defines the change password use case.
type ChangeUserPasswordUseCase struct {
	userRepository *repository.UserRepository
	authUseCase    *AuthUseCase
}

// NewChangePasswordUseCase is a factory function that returns a new instance of the ChangePasswordUseCase.
//
// u: The user repository.
//
// Returns a new instance of the ChangePasswordUseCase.
func NewChangeUserPasswordUseCase(u *repository.UserRepository) *ChangeUserPasswordUseCase {
	return &ChangeUserPasswordUseCase{userRepository: u}
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
func (c *ChangeUserPasswordUseCase) Process(userID uint, form *dto.ChangePasswordForm) (*models.User, *entities.ProcessError) {
	// Get the user by ID
	user, err := c.userRepository.GetUsingID(userID)

	// Check if there is an error
	if err != nil {
		log.Panicln("Error getting user: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while getting the user",
		}
	}

	// Check if the old password is correct
	if !c.authUseCase.VerifyPassword(form.OldPassword, user.Password) {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message: types.FormErrorResponseMsg{
				"old_password": []string{"Old password is incorrect"},
			},
		}
	}

	// Hash the new password
	hashedNewPassword, err := c.authUseCase.HashPassword(form.NewPassword)

	// Check if there is an error
	if err != nil {
		log.Println("Error hashing password: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while hashing the new password",
		}
	}

	// Update the user's password
	user, err = c.userRepository.UpdatePassword(userID, hashedNewPassword)

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

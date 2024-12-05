package usecases

import (
	"fmt"
	"main/core/constants"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"main/pkg/utils"
	"strings"
)

// RegisterUseCase is a struct that defines the register use case.
type RegisterUseCase struct {
	AuthUseCase    *AuthUseCase
	UserRepository *repository.UserRepository
}

// NewRegisterUseCase is a factory function that returns a new instance of the RegisterUseCase.
//
// a: The auth use case.
// u: The user repository.
//
// Returns a new instance of the RegisterUseCase.
func NewRegisterUseCase(a *AuthUseCase, u *repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		AuthUseCase:    a,
		UserRepository: u,
	}
}

// SanitizeUserRegisterForm is a helper function that sanitizes the register input.
//
// form: The register form form dto.
//
// Returns void
func (r RegisterUseCase) SanitizeUserRegisterForm(form *dto.UserRegisterFormDTO) {
	form.Username = strings.TrimSpace(form.Username)
	form.PhoneNumber = strings.TrimSpace(form.PhoneNumber)
}

// ValidateUserRegisterForm is a function that validates the register form.
//
// form: The register form dto.
//
// Returns a map of errors.
func (r RegisterUseCase) ValidateUserRegisterForm(form *dto.UserRegisterFormDTO) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the username is blank
	if utils.IsBlank(form.Username) {
		errs["username"] = append(errs["username"], "Username is required")
	}

	// Check if the username is too short
	if len(form.Username) < constants.MINIMUM_USERNAME_LENGTH {
		errs["username"] = append(errs["username"], fmt.Sprintf("Username must be at least %d characters long", constants.MINIMUM_USERNAME_LENGTH))
	}

	// Check if the phone number is blank
	if utils.IsBlank(form.PhoneNumber) {
		errs["phone_number"] = append(errs["phone_number"], "Phone number is required")
	}

	// Check if the password is blank
	if utils.IsBlank(form.Password) {
		errs["password"] = append(errs["password"], "Password is required")
	}

	// Check if the password is too short
	if len(form.Password) < constants.MINIMUM_PASSWORD_LENGTH {
		errs["password"] = append(errs["password"], fmt.Sprintf("Password must be at least %d characters long", constants.MINIMUM_PASSWORD_LENGTH))
	}

	// Check if the password and confirm password are the same
	if form.Password != form.ConfirmPassword {
		errs["confirm_password"] = append(errs["confirm_password"], "Password and confirm password do not match")
	}

	// Check if there are any errors
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// ProcessUserRegister is a function that processes the register form.
//
// form: The register form dto.
//
// Returns the user and an error message.
func (r RegisterUseCase) ProcessUserRegister(form *dto.UserRegisterFormDTO) *entities.ProcessError {
	// Check if the username is taken
	taken, err := r.UserRepository.IsUsernameTaken(form.Username)

	// Check if there is an error
	if err != nil {
		return &entities.ProcessError{
			Message:     "An error occurred while checking if the username is taken",
			ClientError: false,
		}
	}

	// Check if the username is taken
	if taken {
		return &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"Username is taken"},
			},
			ClientError: true,
		}
	}

	// Check if the phone number is taken
	taken, err = r.UserRepository.IsPhoneNumberTaken(form.PhoneNumber)

	// Check if there is an error
	if err != nil {
		return &entities.ProcessError{
			Message:     "An error occurred while checking if the phone number is taken",
			ClientError: false,
		}
	}

	// Check if the phone number is taken
	if taken {
		return &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"phone_number": []string{"Phone number is taken"},
			},
			ClientError: true,
		}
	}

	return nil
}

// CreateNewUser is a function that creates a new user.
//
// form: The register form dto.
//
// Returns the user and an error message.
func (r *RegisterUseCase) CreateNewUser(form *dto.UserRegisterFormDTO) (*models.User, error) {
	// Hash the password
	hashedPwd, err := r.AuthUseCase.HashPassword(form.Password)

	// Check if there is an error hashing the password
	if err != nil {
		return nil, err
	}

	// Create a new user
	user := models.User{
		Username:    form.Username,
		Password:    hashedPwd,
		PhoneNumber: form.PhoneNumber,
	}

	// Register the user into the database
	err = r.UserRepository.Create(&user)

	// Check if there is an error creating the user
	if err != nil {
		return nil, err
	}

	return &user, nil
}

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
	"strings"
)

// RegisterUseCase is a struct that defines the register use case.
type RegisterUseCase struct {
	userRepository *repository.UserRepository
	authUseCase    *AuthUseCase
}

// NewRegisterUseCase is a factory function that returns a new instance of the RegisterUseCase.
//
// u: The user repository.
//
// Returns a new instance of the RegisterUseCase.
func NewRegisterUseCase(u *repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{userRepository: u}
}

// SanitizeRegisterForm is a helper function that sanitizes the register input.
//
// form: The register form form.
//
// Returns void
func (r RegisterUseCase) SanitizeForm(form *dto.RegisterForm) {
	form.Username = strings.TrimSpace(form.Username)
	form.PhoneNumber = strings.TrimSpace(form.PhoneNumber)
}

// ValidateForm is a function that validates the register form.
//
// form: The register form.
//
// Returns a map of errors.
func (r RegisterUseCase) ValidateForm(form *dto.RegisterForm) types.FormErrorResponseMsg {
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

// Process is a function that processes the register form.
//
// form: The register form.
//
// Returns the user and an error message.
func (r RegisterUseCase) Process(form *dto.RegisterForm) *entities.ProcessError {
	// Check if the username is taken
	taken, err := r.userRepository.IsUsernameTaken(form.Username)

	// Check if there is an error
	if err != nil {
		log.Println("Error checking if username is taken: ", err)

		return &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"An error occurred"},
			},
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
	taken, err = r.userRepository.IsPhoneNumberTaken(form.PhoneNumber)

	// Check if there is an error
	if err != nil {
		log.Println("Error checking if phone number is taken: ", err)

		return &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"phone_number": []string{"An error occurred"},
			},
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
// form: The register form.
//
// Returns the user and an error message.
func (r *RegisterUseCase) CreateNewUser(form *dto.RegisterForm) (*models.User, error) {
	// Hash the password
	hashedPwd, err := r.authUseCase.HashPassword(form.Password)

	// Check if there is an error hashing the password
	if err != nil {
		log.Println("Error hashing the password: ", err)

		return nil, err
	}

	// Create a new user
	newUser := models.User{
		Username:    form.Username,
		Password:    hashedPwd,
		PhoneNumber: form.PhoneNumber,
	}

	// Register the user into the database
	err = r.userRepository.Create(&newUser)

	// Check if there is an error creating the user
	if err != nil {
		log.Println("Error creating the user: ", err)

		return nil, err
	}

	return &newUser, nil
}

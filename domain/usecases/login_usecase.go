package usecases

import (
	"log"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"main/pkg/utils"

	"gorm.io/gorm"
)

// LoginUseCase is a struct that defines the login use case.
type LoginUseCase struct {
	AuthUseCase     *AuthUseCase
	UserRepository  *repository.UserRepository
	VendorRepositoy *repository.VendorRepository
}

// NewLoginUseCase is a factory function that returns a new instance of the LoginUseCase.
//
// a: The auth use case.
// u: The user repository.
//
// Returns a new instance of the LoginUseCase.
func NewLoginUseCase(a *AuthUseCase, u *repository.UserRepository, v *repository.VendorRepository) *LoginUseCase {
	return &LoginUseCase{
		AuthUseCase:     a,
		UserRepository:  u,
		VendorRepositoy: v,
	}
}

// ValidateUserForm is a function that validates the user login form.
//
// form: The login form data.
//
// Returns a map of errors.
func (l LoginUseCase) ValidateUserForm(form *dto.UserLoginForm) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the username is blank
	if utils.IsBlank(form.Username) {
		errs["username"] = append(errs["username"], "Username is required")
	}

	// Check if the password is blank
	if form.Password == "" {
		errs["password"] = append(errs["password"], "Password is required")
	}

	// Return the errors if any
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// ValidateVendorForm is a function that validates the vendor login form.
//
// form: The login form data.
//
// Returns a map of errors.
func (l LoginUseCase) ValidateVendorForm(form *dto.VendorLoginFormDTO) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the email is blank
	if utils.IsBlank(form.Email) {
		errs["email"] = append(errs["email"], "Email is required")
	}

	// Check if the password is blank
	if form.Password == "" {
		errs["password"] = append(errs["password"], "Password is required")
	}

	// Return the errors if any
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Process is a function that processes the user login form.
//
// form: The login form data.
//
// Returns the user object and an error if any.
func (l LoginUseCase) ProcessUser(form *dto.UserLoginForm) (*models.User, *entities.ProcessError) {
	// Check if the username is taken
	user, err := l.UserRepository.GetUsingUsername(form.Username)

	// Check if the user does not exist
	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"Username does not exist"},
			},
			ClientError: true,
		}
	}

	// Return an error if any
	if err != nil {
		log.Println("Error getting the user: ", err)

		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"username": []string{"An error occurred while getting the user"},
			},
			ClientError: false,
		}
	}

	// Check if the password is correct
	if !l.AuthUseCase.VerifyPassword(form.Password, user.Password) {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"password": []string{"Password is incorrect"},
			},
			ClientError: true,
		}
	}

	return user, nil
}

// ProcessVendor is a function that processes the vendor login form.
//
// form: The login form data.
//
// Returns the vendor object and an error if any.
func (l LoginUseCase) ProcessVendor(form *dto.VendorLoginFormDTO) (*models.Vendor, *entities.ProcessError) {
	// Check if the email is available
	vendor, err := l.VendorRepositoy.GetUsingEmail(form.Email)

	// Check if the vendor does not exist
	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"email": []string{"Email does not exist"},
			},
			ClientError: true,
		}
	}

	// Return an error if any
	if err != nil {
		log.Println("Error getting the vendor: ", err)

		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"email": []string{"An error occurred while getting the vendor"},
			},
			ClientError: false,
		}
	}

	// Check if the password is correct
	if !l.AuthUseCase.VerifyPassword(form.Password, vendor.Password) {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"password": []string{"Password is incorrect"},
			},
			ClientError: true,
		}
	}

	return vendor, nil
}

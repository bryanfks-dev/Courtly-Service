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
	AuthUseCase      *AuthUseCase
	UserRepository   *repository.UserRepository
	VendorRepository *repository.VendorRepository
}

// VerifyPassword is a method of the VerifyPasswordUseCase that verifies a user's password.
//
// a: The AuthUseCase instance.
// u: The UserRepository instance.
// v: The VendorRepository instance.
//
// Retuns instance of the VerifyPasswordUseCase.
func NewVerifyPasswordUseCase(a *AuthUseCase, u *repository.UserRepository, v *repository.VendorRepository) *VerifyPasswordUseCase {
	return &VerifyPasswordUseCase{
		AuthUseCase:      a,
		UserRepository:   u,
		VendorRepository: v,
	}
}

// ValidateForm is a function that validates the verify password form.
//
// form: The verify password form dto.
//
// Returns a map of errors.
func (*VerifyPasswordUseCase) ValidateForm(form *dto.VerifyPasswordFormDTO) types.FormErrorResponseMsg {
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

// ProcessUser is a function that processes the user verify password form.
//
// form: The verify password form dto.
// userID: The user ID.
//
// Returns the user object and an error if any.
func (v *VerifyPasswordUseCase) ProcessUser(form *dto.VerifyPasswordFormDTO, token *jwt.Token) (*models.User, *entities.ProcessError) {
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

// ProcessVendor is a function that processes the vendor verify password form.
//
// form: The verify password form dto.
// vendorID: The vendor ID.
//
// Returns the user object and an error if any.
func (v *VerifyPasswordUseCase) ProcessVendor(form *dto.VerifyPasswordFormDTO, token *jwt.Token) (*models.Vendor, *entities.ProcessError) {
	// Decode the token
	claims := v.AuthUseCase.DecodeToken(token)

	// Get the vendor by the vendor ID
	vendor, err := v.VendorRepository.GetUsingID(claims.Id)

	// Check if there is an error
	if err != nil {
		log.Println("Error getting vendor using ID: ", err)

		return nil, &entities.ProcessError{
			Message:     "An error occurred while getting the vendor",
			ClientError: false,
		}
	}

	// Check if the password is correct
	if !v.AuthUseCase.VerifyPassword(form.Password, vendor.Password) {
		return nil, &entities.ProcessError{
			Message: types.FormErrorResponseMsg{
				"password": []string{"Incorrect password"},
			},
			ClientError: true,
		}
	}

	return vendor, nil
}

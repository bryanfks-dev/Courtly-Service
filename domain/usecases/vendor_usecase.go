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

	"github.com/golang-jwt/jwt/v5"
)

// VendorUseCase is a struct that defines the vendor use case.
type VendorUseCase struct {
	AuthUseCase      *AuthUseCase
	VendorRepository *repository.VendorRepository
}

// NewVendorUseCase is a factory function that returns a new instance of the VendorUseCase.
//
// a: The auth use case.
// v: The vendor repository.
//
// Returns a new instance of the VendorUseCase.
func NewVendorUseCase(a *AuthUseCase, v *repository.VendorRepository) *VendorUseCase {
	return &VendorUseCase{
		AuthUseCase:      a,
		VendorRepository: v,
	}
}

// GetCurrentVendor is a function that returns the current vendor.
//
// token: The token.
//
// Returns the current vendor and an error if any.
func (v *VendorUseCase) GetCurrentVendor(token *jwt.Token) (*models.Vendor, *entities.ProcessError) {
	// Get the token claims
	claims := v.AuthUseCase.DecodeToken(token)

	// Get the vendor by ID
	vendor, err := v.VendorRepository.GetUsingID(claims.Id)

	// Check if there is an error
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while getting the vendor",
		}
	}

	return vendor, nil
}

// ValidateChangePasswordForm is a function that validates the change password form.
//
// form: The change password form dto.
//
// Returns a map of errors.
func (v *VendorUseCase) ValidateChangePasswordForm(form *dto.ChangePasswordFormDTO) types.FormErrorResponseMsg {
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

	// Check if the new password is less than the minimum password length
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
// token: The vendor token.
// form: The change password form dto.
//
// Returns an error if any.
func (v *VendorUseCase) ProcessChangePassword(token *jwt.Token, form *dto.ChangePasswordFormDTO) *entities.ProcessError {
	// Get the vendor ID from the token
	claims := v.AuthUseCase.DecodeToken(token)

	// Get the vendor by ID
	vendor, err := v.VendorRepository.GetUsingID(claims.Id)

	// Check if there is an error
	if err != nil {
		return &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while getting the vendor",
		}
	}

	// Check if the old password is correct
	if !v.AuthUseCase.VerifyPassword(form.OldPassword, vendor.Password) {
		return &entities.ProcessError{
			ClientError: true,
			Message: types.FormErrorResponseMsg{
				"old_password": []string{"Old password is incorrect"},
			},
		}
	}

	// Hash the new password
	hashedNewPassword, err := v.AuthUseCase.HashPassword(form.NewPassword)

	// Check if there is an error
	if err != nil {
		return &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while hashing the new password",
		}
	}

	// Update the vendor's password
	err = v.VendorRepository.UpdatePassword(claims.Id, hashedNewPassword)

	// Check if there is an error
	if err != nil {
		return &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while updating the vendor's password",
		}
	}

	return nil
}

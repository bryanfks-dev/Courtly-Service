package usecases

import (
	"log"
	"main/data/models"
	"main/domain/entities"
	"main/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

// GetVendorUsingID is a function that returns a vendor using the given ID.
//
// vendorID: The vendor ID.
//
// Returns the vendor and an error if any.
func (v *VendorUseCase) GetVendorUsingID(vendorID uint) (*models.Vendor, *entities.ProcessError) {
	// Get the vendor from the database
	vendor, err := v.VendorRepository.GetUsingID(vendorID)

	// Return an error if the vendor is not found
	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			Message:     "Vendor not found",
			ClientError: true,
		}
	}

	// Return an error if any
	if err != nil {
		log.Println("Failed to get current vendor: ", err)

		return nil, &entities.ProcessError{
			Message:     "Failed to get current vendor",
			ClientError: false,
		}
	}

	return vendor, nil
}

// GetCurrentVendor is a function that returns the current vendor.
//
// token: The token.
//
// Returns the current vendor and an error if any.
func (v *VendorUseCase) GetCurrentVendor(token *jwt.Token) (*models.Vendor, *entities.ProcessError) {
	// Get the token claims
	claims := v.AuthUseCase.DecodeToken(token)

	return v.GetVendorUsingID(claims.Id)
}

package usecases

import (
	"log"
	"main/data/models"
	"main/internal/dto"
	"main/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// CourtUseCase is a struct that defines the use case for the court entity.
type CourtUseCase struct {
	AuthUseCase     *AuthUseCase
	CourtRepository *repository.CourtRepository
}

// NewCourtUseCase is a factory function that returns a new instance of the CourtUseCase struct.
//
// a: The auth use case.
// c: The court repository.
//
// Returns a new instance of the CourtUseCase.
func NewCourtUseCase(a *AuthUseCase, c *repository.CourtRepository) *CourtUseCase {
	return &CourtUseCase{
		AuthUseCase:     a,
		CourtRepository: c,
	}
}

// GetVendorCourtsUsingType is a function that returns the vendor courts using the court type.
//
// vendorID: The vendor ID.
//
// Returns the vendor courts and an error if any.
func (c *CourtUseCase) GetVendorCourtsUsingType(vendorID uint, courtType string) (*[]models.Court, error) {
	// Get the vendor courts using the court type
	courts, err := c.CourtRepository.GetVendorCourtsUsingType(vendorID, courtType)

	// Return an error if any and ignore if the vendor has no courts
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get vendor courts using type: ", err)

		return nil, err
	}

	return courts, nil
}

// GetCurrentVendorCourts is a function that returns the current vendor courts.
//
// token: The token.
//
// Returns the vendor courts and an error if any.
func (c *CourtUseCase) GetCurrentVendorCourtsUsingType(token *jwt.Token, courtType string) (*[]models.Court, error) {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the vendor courts
	return c.GetVendorCourtsUsingType(claims.Id, courtType)
}

// ConvertCourtModelsToDTOs is a function that converts a list of court models to a list of court DTOs.
//
// courts: The list of court models.
//
// Returns a list of court DTOs.
func (c *CourtUseCase) ConvertCourtModelsToDTOs(courts *[]models.Court) *[]dto.VendorCourt {
	// Create a new list of vendor courts
	var vendorCourts []dto.VendorCourt

	// Convert the court models to court DTOs
	for _, court := range *courts {
		vendorCourts = append(vendorCourts, *dto.VendorCourt{}.FromModel(court))
	}

	return &vendorCourts
}

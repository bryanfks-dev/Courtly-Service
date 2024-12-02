package usecases

import (
	"log"
	"main/core/enums"
	"main/data/models"
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

// GetCourts is a function that returns the courts.
//
// Returns the courts and an error if any.
func (c *CourtUseCase) GetCourts() (*[]models.Court, error) {
	// Get the courts
	courts, err := c.CourtRepository.GetAll()

	// Return an error if any
	if err != nil {
		log.Println("Failed to get courts: ", err)

		return nil, err
	}

	return courts, nil
}

// GetCourtUsingID is a function that returns the court using the court ID.
//
// courtID: The court ID.
//
// Returns the court and an error if any.
func (c *CourtUseCase) GetCourtUsingID(courtID uint) (*models.Court, error) {
	// Get the courts
	court, err := c.CourtRepository.GetUsingID(courtID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get courts: ", err)

		return nil, err
	}

	return court, nil
}

// ValidateCourtType is a function that validates the court type.
//
// courtType: The court type.
//
// Returns true if the court type is valid.
func (c *CourtUseCase) ValidateCourtType(courtType string) bool {
	return enums.InCourtType(courtType)
}

// GetCourtsUsingCourtType is a function that returns the courts using the court type.
//
// courtType: The court type.
//
// Returns the courts and an error if any.
func (c *CourtUseCase) GetCourtsUsingCourtType(courtType string) (*[]models.Court, error) {
	// Get the courts
	courts, err := c.CourtRepository.GetUsingCourtType(courtType)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get courts using type ", courtType, ": ", err)

		return nil, err
	}

	return courts, nil
}

// GetVendorCourtsUsingCourtType is a function that returns the vendor courts 
// using the court type.
//
// vendorID: The vendor ID.
//
// Returns the vendor courts and an error if any.
func (c *CourtUseCase) GetVendorCourtsUsingCourtType(vendorID uint, courtType string) (*[]models.Court, error) {
	// Get the vendor courts using the court type
	courts, err := c.CourtRepository.GetUsingVendorIDCourtType(vendorID, courtType)

	// Return an error if any and ignore if the vendor has no courts
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get vendor courts using court type: ", err)

		return nil, err
	}

	return courts, nil
}

// GetCurrentVendorCourtsUsingCourtType is a function that returns the current vendor courts
// with the given court type.
//
// token: The token.
//
// Returns the vendor courts and an error if any.
func (c *CourtUseCase) GetCurrentVendorCourtsUsingCourtType(token *jwt.Token, courtType string) (*[]models.Court, error) {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the vendor courts
	return c.GetVendorCourtsUsingCourtType(claims.Id, courtType)
}

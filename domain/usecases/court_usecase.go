package usecases

import (
	"encoding/base64"
	"fmt"
	"log"
	"main/core/constants"
	"main/core/enums"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"main/pkg/utils"
	"os"
	"strconv"

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
// courtType: The court type.
//
// Returns the vendor courts and an error if any.
func (c *CourtUseCase) GetCurrentVendorCourtsUsingCourtType(token *jwt.Token, courtType string) (*[]models.Court, error) {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the vendor courts
	return c.GetVendorCourtsUsingCourtType(claims.Id, courtType)
}

// ValidateCreateNewCourtForm is a function that validates the create new court form.
//
// form: The CreateNewCourtForm dto.
//
// Returns the form error response message.
func (c *CourtUseCase) ValidateCreateNewCourtForm(form *dto.CreateNewCourtFormDTO) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the price per hour is less than or equal to 0
	if form.PricePerHour <= 0 {
		errs["price_per_hour"] = append(errs["price_per_hour"], "Price per hour must be greater than 0")
	}

	// Check if the court image is blank
	if utils.IsBlank(form.CourtImage) {
		errs["courts_image"] = append(errs["courts_image"], "Court image is required")
	}

	// Return the errors if any
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// GetVendorNewestCourtUsingCourtType is a function that returns the vendor newest court
// using the court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the court and an error if any.
func (c *CourtUseCase) GetVendorNewestCourtUsingCourtType(vendorID uint, courtType string) (*models.Court, error) {
	// Get the vendor newest court using the court type
	courts, err := c.CourtRepository.GetNewestUsingVendorIDCourtType(vendorID, courtType)

	// Return an error if any
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get vendor newest court using court type: ", err)

		return nil, err
	}

	return courts, nil
}

// CreateNewCourt is a function that creates a new court.
//
// token: The token.
// courtType: The court type.
// form: The CreateNewCourtForm dto.
//
// Returns an error if any.
func (c *CourtUseCase) CreateNewCourt(token *jwt.Token, courtType string, form *dto.CreateNewCourtFormDTO) (*models.Court, *entities.ProcessError) {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the vendor newest court using the court type
	court, err := c.GetVendorNewestCourtUsingCourtType(claims.Id, courtType)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while getting current vendor newest court",
		}
	}

	// Check if the court is nil
	if court != nil {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "Vendor already has a court in this court type, use POST /vendor/me/courts/types/:type instead",
		}
	}

	// Decode the image
	fileBytes, err := base64.StdEncoding.DecodeString(form.CourtImage)

	// Return an error if any
	if err != nil {
		log.Println("Failed to decode court image: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while decoding court image",
		}
	}

	// Create the court image name
	courtImageName := fmt.Sprintf("court_%s_%s.jpg", claims.ID, courtType)

	// Write the image to a file
	err = os.WriteFile(fmt.Sprintf("%s/%s", constants.PATH_TO_COURT_IMAGES, courtImageName), fileBytes, 0644)

	// Return an error if any
	if err != nil {
		log.Println("Failed to save court image: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while saving court image",
		}
	}

	// Create a new court object
	newCourt := &models.Court{
		VendorID: claims.Id,
		CourtType: models.CourtType{
			Type: courtType,
		},
		Name:  "Court 1",
		Price: form.PricePerHour,
		Image: courtImageName,
	}

	// Return an error if any
	err = c.CourtRepository.Create(newCourt)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while creating a new court",
		}
	}

	return court, nil
}

// AddCourt is a function that adds a new court.
//
// token: The token.
// courtType: The court type.
//
// Returns the court and an error if any.
func (c *CourtUseCase) AddCourt(token *jwt.Token, courtType string) (*models.Court, *entities.ProcessError) {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the vendor newest court using the court type
	court, err := c.GetVendorNewestCourtUsingCourtType(claims.Id, courtType)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while getting current vendor newest court",
		}
	}

	// Check if the court is nil
	if court == nil {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "Vendor doesn't have any court in this court type, use POST /vendor/me/courts/types/:type/new instead",
		}
	}

	// Create new court name
	courtName := "Court "

	// Append number into new court name
	courtNumber, err := strconv.Atoi(court.Name[len(courtName):])

	// Return an error if any
	if err != nil {
		log.Println("Failed to convert court number: ", err)

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while converting court number",
		}
	}

	// Append new court number
	courtName += strconv.Itoa(courtNumber + 1)

	// Create a new court object
	newCourt := &models.Court{
		VendorID:  claims.Id,
		CourtType: court.CourtType,
		Name:      courtName,
		Price:     court.Price,
		Image:     court.Image,
	}

	// Create the new court
	err = c.CourtRepository.Create(newCourt)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while creating a new court",
		}
	}

	return newCourt, nil
}

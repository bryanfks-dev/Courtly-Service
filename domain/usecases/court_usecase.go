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
	return c.CourtRepository.GetAll()
}

// GetCourtUsingID is a function that returns the court using the court ID.
//
// courtID: The court ID.
//
// Returns the court and an error if any.
func (c *CourtUseCase) GetCourtUsingID(courtID uint) (*models.Court, error) {
	// Get the courts
	return c.CourtRepository.GetUsingID(courtID)
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
	return c.CourtRepository.GetUsingCourtType(courtType)
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
	return c.CourtRepository.GetUsingVendorIDCourtType(claims.Id, courtType)
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

	// Check if the court already exist
	exist, err := c.CourtRepository.CheckExistUsingVendorIDCourtType(claims.Id, courtType)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while checking court exist in this court type",
		}
	}

	// Check if the court is nil
	if exist {
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
	court := &models.Court{
		VendorID: claims.Id,
		CourtType: models.CourtType{
			Type: courtType,
		},
		Name:  "Court 1",
		Price: form.PricePerHour,
		Image: courtImageName,
	}

	// Return an error if any
	err = c.CourtRepository.Create(court)

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

	// Check if the court already exist
	exist, err := c.CourtRepository.CheckExistUsingVendorIDCourtType(claims.Id, courtType)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while checking current vendor court exist in this court type",
		}
	}

	// Check if the court is nil
	if !exist {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "Vendor doesn't have any court in this court type, use POST /vendor/me/courts/types/:type/new instead",
		}
	}

	// Get the newest court using vendor ID and court type
	court, err := c.CourtRepository.GetNewestUsingVendorIDCourtType(claims.Id, courtType)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occured while getting the newest court",
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

// GetCurrentVendorCourtStats is a function that returns the current vendor court stats.
//
// token: The token.
//
// Returns the vendor court count as map and an error if any.
func (c *CourtUseCase) GetCurrentVendorCourtStats(token *jwt.Token) (*types.CourtCountsMap, error) {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Get the court counts
	return c.CourtRepository.GetCountsUsingVendorID(claims.Id)
}

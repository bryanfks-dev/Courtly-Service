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
	"sort"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// CourtUseCase is a struct that defines the use case for the court entity.
type CourtUseCase struct {
	AuthUseCase      *AuthUseCase
	CourtRepository  *repository.CourtRepository
	ReviewRepository *repository.ReviewRepository
}

// NewCourtUseCase is a factory function that returns a new instance of the CourtUseCase struct.
//
// a: The auth use case.
// c: The court repository.
// r: The review repository.
//
// Returns a new instance of the CourtUseCase.
func NewCourtUseCase(a *AuthUseCase, c *repository.CourtRepository, r *repository.ReviewRepository) *CourtUseCase {
	return &CourtUseCase{
		AuthUseCase:      a,
		CourtRepository:  c,
		ReviewRepository: r,
	}
}

// GetCourts is a function that returns the courts.
//
// courtType: The court type.
// search: The search query.
//
// Returns the courts and an error if any.
func (c *CourtUseCase) GetCourts(courtType *string, search *string) (*[]types.CourtMap, error) {
	// Create an empty courts slice and an error
	var (
		courts *[]models.Court
		err    error
	)

	// Get the courts
	if (courtType == nil || utils.IsBlank(*courtType)) && (search == nil || utils.IsBlank(*search)) {
		courts, err = c.CourtRepository.Get()
	} else if (courtType != nil && !utils.IsBlank(*courtType)) && (search == nil || utils.IsBlank(*search)) {
		courts, err = c.CourtRepository.GetUsingCourtType(*courtType)
	} else if (courtType == nil || utils.IsBlank(*courtType)) && (search != nil && !utils.IsBlank(*search)) {
		courts, err = c.CourtRepository.GetUsingVendorName(*search)
	} else {
		courts, err = c.CourtRepository.GetUsingCourtTypeVendorName(*courtType, *search)
	}

	// Return an error if any
	if err != nil {
		return nil, err
	}

	// Create a new court maps slice
	courtMaps := make([]types.CourtMap, len(*courts))

	// Loop through the courts
	for i, court := range *courts {
		// Get the court average rating
		totalRating, err := c.ReviewRepository.GetAvgRatingUsingCourtTypeVendorID(court.CourtType.Type, court.VendorID)

		// Return an error if any
		if err != nil {
			return nil, err
		}

		// Append the court map
		courtMaps[i] = types.CourtMap{
			"court":        court,
			"total_rating": totalRating,
		}
	}

	// Sort the courts based by total rating
	sort.Slice(*courts, func(i, j int) bool {
		return courtMaps[i].GetTotalRating() > courtMaps[j].GetTotalRating()
	})

	return &courtMaps, nil
}

// GetVendorCourtsUsingCourtType is a function that returns the vendor courts with the given court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the vendor courts map and an error if any.
func (c *CourtUseCase) GetVendorCourtsUsingCourtType(vendorID uint, courtType string) (*[]types.CourtMap, error) {
	// Get the courts
	courts, err := c.CourtRepository.GetUsingVendorIDCourtType(vendorID, courtType)

	// Return an error if any
	if err != nil {
		return nil, err
	}

	// Create a new court maps slice
	courtMaps := make([]types.CourtMap, len(*courts))

	// Get average rating for the court type
	totalRating, err := c.ReviewRepository.GetAvgRatingUsingCourtTypeVendorID(courtType, vendorID)

	// Loop through the courts
	for i, court := range *courts {
		// Return an error if any
		if err != nil {
			return nil, err
		}

		// Append the court map
		courtMaps[i] = types.CourtMap{
			"court":        court,
			"total_rating": totalRating,
		}
	}

	return &courtMaps, nil
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

	// Get the court type ID
	courtTypeId := enums.GetCourtTypeID(courtType)

	// Create the court image name
	courtImageName := fmt.Sprintf("court_%d_%d.jpg", claims.Id, courtTypeId)

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
		VendorID:    claims.Id,
		CourtTypeID: uint(courtTypeId),
		Name:        "Court 1",
		Price:       form.PricePerHour,
		Image:       courtImageName,
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

// ValidateUpdateCourtForm is a function to validate the update court form.
//
// form: The update court form dto
//
// Returns a form error response message
func (c *CourtUseCase) ValidateUpdateCourtForm(form *dto.UpdateCourtFormDTO) types.FormErrorResponseMsg {
	// Make an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if the price per hour is less than or equal to 0
	if form.PricePerHour <= 0 {
		errs["price_per_hour"] = append(errs["price_per_hour"], "Price per hour must be greater than 0")
	}

	// Check if error is exists
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// UpdateCourtUsingCourtType is a function to update court using the given court type.
//
// token: The jwt token
// courtType: The court type
//
// Returns error if any
func (c *CourtUseCase) UpdateCourtUsingCourtType(token *jwt.Token, courtType string, form *dto.UpdateCourtFormDTO) error {
	// Get the token claims
	claims := c.AuthUseCase.DecodeToken(token)

	// Update the court
	return c.CourtRepository.UpdateUsingVendorIDCourtType(claims.Id, courtType, form.PricePerHour)
}

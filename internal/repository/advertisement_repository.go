package repository

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"
)

// AdvertisementRepository is the repository for the advertisement.
type AdvertisementRepository struct{}

// NewAdvertisementRepository is a constructor for the AdvertisementRepository.
//
// Returns a new instance of AdvertisementRepository.
func NewAdvertisementRepository() *AdvertisementRepository {
	return &AdvertisementRepository{}
}

// Create is a method to create an advertisement.
//
// ad: The advertisement to create.
//
// Returns an error if any.
func (*AdvertisementRepository) Create(ad *models.Advertisement) error {
	// Create the advertisement
	err := mysql.Conn.Create(ad).Error

	// Return error if any
	if err != nil {
		log.Println("Error create advertisement: " + err.Error())

		return err
	}

	return nil
}

// GetAll is a method to get all advertisements.
//
// Returns a slice of advertisements and an error if any.
func (*AdvertisementRepository) GetAll() (*[]models.Advertisement, error) {
	// Create a variable to store advertisements
	var ads []models.Advertisement

	// Get all advertisements
	err := mysql.Conn.Preload("Vendor").Preload("CourtType").Find(&ads).Error

	// Log the error if any
	if err != nil {
		log.Println("Error get all advertisements: " + err.Error())

		return nil, err
	}

	// Return the advertisements and error if any
	return &ads, err
}

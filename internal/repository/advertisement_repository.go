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

// GetAll is a method to get all advertisements.
//
// Returns a slice of advertisements and an error if any.
func (*AdvertisementRepository) GetAll() (*[]models.Advertisement, error) {
	// Create a variable to store advertisements
	var ads []models.Advertisement

	// Get all advertisements
	err := mysql.Conn.Find(&ads).Error

	// Log the error if any
	if err != nil {
		log.Println("Error get all advertisements: " + err.Error())

		return nil, err
	}

	// Return the advertisements and error if any
	return &ads, err
}

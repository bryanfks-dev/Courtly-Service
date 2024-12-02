package repository

import (
	"main/data/models"
	"main/internal/providers/mysql"
)

// ReviewRepository is a struct that defines the ReviewRepository
type ReviewRepository struct{}

// NewReviewRepository is a factory function that returns a new instance of the ReviewRepository.
//
// Returns a new instance of the ReviewRepository.
func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{}
}

// GetUsingVendorIDCourtType is a function that returns the reviews using the vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the reviews and an error if any.
func (*ReviewRepository) GetUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Review, error) {
	// reviews is a slice of maps containing the reviews of the court
	var reviews []models.Review

	// Get the reviews using the vendor ID and court type
	err := mysql.Conn.Preload("CourtType", "type = ?", courtType).Where("vendor_id = ?", vendorID).Find(&reviews).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &reviews, err
}

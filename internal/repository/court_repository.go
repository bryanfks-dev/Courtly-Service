package repository

import (
	"main/data/models"
	"main/internal/providers/mysql"
)

// CourtRepository is a struct that defines the court repository.
type CourtRepository struct{}

// NewCourtRepository is a factory function that returns a new instance of the court repository.
//
// Returns a new instance of the court repository.
func NewCourtRepository() *CourtRepository {
	return &CourtRepository{}
}

// GetVendorCourtsUsingType is a function that returns the vendor courts using the court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the vendor courts and an error if any.
func (*CourtRepository) GetVendorCourtsUsingType(vendorID uint, courtType string) (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts by vendor ID and court type
	err := mysql.Conn.Where("vendor_id = ? AND court_type = ?", vendorID, courtType).Find(&courts).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &courts, nil
}

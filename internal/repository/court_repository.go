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

// GetAll is a function that returns all the courts.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetAll() (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts
	err := mysql.Conn.Find(&courts).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &courts, nil
}

// GetUsingID is a function that returns the courts by ID.
//
// courtID: The court ID.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetUsingID(courtID uint) (*models.Court, error) {
	// Create a new court object
	var courts models.Court

	// Get the courts by ID
	err := mysql.Conn.Where("id = ?", courtID).Find(&courts).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &courts, nil
}

// GetUsingType is a function that returns the courts using the court type.
//
// courtType: The court type.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetAllUsingType(courtType string) (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts
	err := mysql.Conn.Joins("JOIN court_types").Where("type = ?", courtType).Find(&courts).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &courts, nil
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

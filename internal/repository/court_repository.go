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

// GetCourts is a function that returns the courts.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetCourts() (*[]models.Court, error) {
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

// GetCourtsUsingType is a function that returns the courts using the court type.
//
// courtType: The court type.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetCourtsUsingType(courtType string) (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts
	err := mysql.Conn.Model(&models.Court{}).Joins("JOIN court_types").Where("type = ?", courtType).Find(&courts).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &courts, nil
}

// GetVendorCourtTypes is a function that returns the vendor court types.
//
// vendorID: The vendor ID.
//
// Returns the vendor court types and an error if any.
func (*CourtRepository) GetVendorCourtTypes(vendorID uint) (*[]models.CourtType, error) {
	// Create a new court type object
	var courtTypes []models.CourtType

	// Get the court types by vendor ID
	err := mysql.Conn.Model(&models.Court{}).Distinct("court_types.type").Joins("JOIN court_types").Where("vendor_id = ?", vendorID).Find(&courtTypes).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &courtTypes, nil
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

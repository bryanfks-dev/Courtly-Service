package repository

import (
	"log"
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

// Create is a function that creates a new court.
//
// court: The court object.
//
// Returns an error if any.
func (*CourtRepository) Create(court *models.Court) error {
	// Create the court
	err := mysql.Conn.Create(court).Error

	if err != nil {
		log.Println("Error creating court: " + err.Error())

		return err
	}

	return nil
}

// GetNewestUsingVendorIDCourtType is a function that returns the newest court by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the court and an error if any.
func (*CourtRepository) GetNewestUsingVendorIDCourtType(vendorID uint, courtType string) (*models.Court, error) {
	// Create a new court object
	var court models.Court

	// Get the courts by vendor ID and court type
	err := mysql.Conn.Preload("CourtType", "type = ?", courtType).Where("vendor_id = ?", vendorID).Order("created_at desc").First(&court).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting newest court using vendor ID and court type: " + err.Error())

		return nil, err
	}

	return &court, nil
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
		log.Println("Error getting all courts: " + err.Error())

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
		log.Println("Error getting court using id: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// GetUsingCourtType is a function that returns the courts using the court type.
//
// courtType: The court type.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetUsingCourtType(courtType string) (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts
	err := mysql.Conn.Preload("CourtType", "type = ?", courtType).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting courts using court type: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// GetUsingVendorIDCourtType is a function that returns the courts by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the vendor courts and an error if any.
func (*CourtRepository) GetUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts by vendor ID and court type
	err := mysql.Conn.Preload("CourtType", "type = ?", courtType).Where("vendor_id = ?", vendorID).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting courts using vendor id and court type: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// CheckExistsUsingVendorIDCourtType is a function that checks if the courts exist by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns a boolean and an error if any.
func (*CourtRepository) CheckExistUsingVendorIDCourtType(vendorID uint, courtType string) (bool, error) {
	// count is the number of courts
	var count int64

	// Get the courts by vendor ID and court type
	err := mysql.Conn.Model(&models.Court{}).Preload("CourtType", "type = ?", courtType).Where("vendor_id = ?", vendorID).Limit(1).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking court exist using vendor id and court type: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

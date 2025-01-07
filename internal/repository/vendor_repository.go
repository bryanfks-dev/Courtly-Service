package repository

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"
)

// VendorRepository is a struct that defines the vendor repository.
type VendorRepository struct{}

// NewVendorRepository is a factory function that returns a new instance of the vendor repository.
//
// Returns a new instance of the vendor repository.
func NewVendorRepository() *VendorRepository {
	return &VendorRepository{}
}

// Create is a function that creates a new vendor.
//
// vendor: The vendor object.
//
// Returns an error if any.
func (*VendorRepository) Create(vendor *models.Vendor) error {
	// Create the vendor
	err := mysql.Conn.Create(vendor).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to create vendor: " + err.Error())

		return err
	}

	return nil
}

// GetUsingID is a function that returns a vendor by ID.
//
// vendorID: The vendor ID.
//
// Returns the vendor object and an error if any.
func (*VendorRepository) GetUsingID(vendorID uint) (*models.Vendor, error) {
	// Create a new vendor object
	var vendor models.Vendor

	// Get the vendor by ID
	err := mysql.Conn.First(&vendor, "id = ?", vendorID).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to get vendor using id: " + err.Error())

		return nil, err
	}

	return &vendor, err
}

// IsEmailTaken is a function that checks if an email is taken.
//
// email: The email.
//
// Returns a boolean and an error if any.
func (*VendorRepository) IsEmailTaken(email string) (bool, error) {
	// Create a counter variable
	var count int64

	// Check if the email is taken
	err := mysql.Conn.Model(&models.Vendor{}).Where("email = ?", email).Limit(1).Count(&count).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to check if email is taken: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

// GetUsingEmail is a function that returns a vendor by email.
//
// email: The email.
//
// Returns the vendor object and an error if any.
func (*VendorRepository) GetUsingEmail(email string) (*models.Vendor, error) {
	// Create a new vendor object
	var vendor models.Vendor

	// Get the vendor by email
	err := mysql.Conn.First(&vendor, "email = ?", email).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to get vendor using email: " + err.Error())

		return nil, err
	}

	return &vendor, err
}

// UpdatePassword is a function that updates a vendor's password.
//
// vendorID: The vendor ID.
// hashedNewPassword: The hashed new password.
//
// Returns an error if any.
func (*VendorRepository) UpdatePassword(vendorID uint, hashedNewPassword string) error {
	// Update the vendor's password
	err := mysql.Conn.Model(&models.Vendor{}).Where("id = ?", vendorID).Update("password", hashedNewPassword).Error

	// Check if there is an error
	if err != nil {
		log.Println("Failed to update vendor password: " + err.Error())

		return err
	}

	return nil
}

package repository

import (
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
	return mysql.Conn.Create(vendor).Error
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

	return &vendor, err
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

	return &vendor, err
}

// UpdatePassword is a function that updates a vendor's password.
//
// vendorID: The vendor ID.
// hashedNewPassword: The hashed new password.
//
// Returns an error if any.
func (*VendorRepository) UpdatePassword(vendorID uint, hashedNewPassword string) (*models.Vendor, error) {
	var vendor models.Vendor

	// Get the vendor by ID
	err := mysql.Conn.First(&vendor, "id = ?", vendorID).Error

	// Check if there is an error
	if err != nil {
		return nil, err
	}

	// Update the vendor's password
	err = mysql.Conn.Model(&vendor).Update("password", hashedNewPassword).Error

	// Check if there is an error
	if err != nil {
		return nil, err
	}

	return &vendor, err
}

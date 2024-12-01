package repository

import (
	"main/data/models"
	"main/internal/providers/mysql"
)

// BookingRepository is a struct that defines the BookingRepository
type BookingRepository struct{}

// NewBookingRepository is a function that returns a new BookingRepository
//
// Returns a pointer to the BookingRepository struct
func NewBookingRepository() *BookingRepository {
	return &BookingRepository{}
}

// GetByUserID is a method that returns the bookings by the given user ID.
//
// userID: The ID of the user.
//
// Returns the bookings and an error if any.
func (*BookingRepository) GetByUserID(userID uint) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	// Get the bookings from the database
	err := mysql.Conn.Where("user_id = ?", userID).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &bookings, nil
}

// GetByVendorID is a method that returns the bookings by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the bookings and an error if any.
func (*BookingRepository) GetByVendorID(vendorID uint) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	// Get the bookings from the database
	err := mysql.Conn.Where("vendor_id = ?", vendorID).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &bookings, nil
}

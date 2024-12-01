package repository

import (
	"main/core/enums"
	"main/data/models"
	"main/internal/providers/mysql"
	"time"
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
	err := mysql.Conn.Joins("JOIN orders").Where("vendor_id = ? AND orders.status = ?", vendorID, enums.Success.Label()).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &bookings, nil
}

// GetTotalByVendorID is a method that returns the total bookings by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the total bookings and an error if any.
func (*BookingRepository) GetTotalByVendorID(vendorID uint) (int64, error) {
	// count is a placeholder for the count
	var count int64

	// Get the bookings from the database
	err := mysql.Conn.Model(&models.Booking{}).Joins("JOIN orders").Where("vendor_id = ? AND orders.status = ?", vendorID, enums.Success.Label()).Count(&count).Error

	// Return an error if any
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetTotalTodayByVendorID is a method that returns the total bookings today by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the total bookings today and an error if any.
func (*BookingRepository) GetTotalTodayByVendorID(vendorID uint) (int64, error) {
	// count is a placeholder for the count
	var count int64

	// Get the current date
	today := time.Now().Truncate(24 * time.Hour)

	// Get the bookings from the database
	err := mysql.Conn.Model(&models.Booking{}).Joins("JOIN orders").Where("vendor_id = ? AND orders.status = ? AND date >= ? AND date < ?", vendorID, enums.Success.Label(), today, today.Add(24*time.Hour)).Count(&count).Error

	// Return an error if any
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetNLatestByVendorID is a method that returns the n latest bookings by the given vendor ID.
//
// vendorID: The ID of the vendor.
// n: The number of bookings to return.
//
// Returns the n latest bookings and an error if any.
func (*BookingRepository) GetNLatestByVendorID(vendorID uint, n int) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	// Get the bookings from the database
	err := mysql.Conn.Joins("JOIN orders").Where("vendor_id = ? AND orders.status = ?", vendorID, enums.Success.Label()).Order("date desc").Limit(n).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return &bookings, nil
}

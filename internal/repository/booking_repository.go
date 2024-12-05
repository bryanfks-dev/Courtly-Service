package repository

import (
	"log"
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

// GetUsingUserID is a method that returns the bookings by the given user ID.
//
// userID: The ID of the user.
//
// Returns the bookings and an error if any.
func (*BookingRepository) GetUsingUserID(userID uint) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	// Get the bookings from the database
	err := mysql.Conn.Where("user_id = ?", userID).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting bookings using user id: " + err.Error())

		return nil, err
	}

	return &bookings, nil
}

// GetUsingVendorID is a method that returns the bookings by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the bookings and an error if any.
func (*BookingRepository) GetUsingVendorID(vendorID uint) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	// Get the bookings from the database
	err := mysql.Conn.Model(&models.Booking{}).Preload("Order", "status = ?", enums.Success.Label()).Where("vendor_id = ?", vendorID).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting bookings using vendor id: " + err.Error())

		return nil, err
	}

	return &bookings, nil
}

// GetTotalUsingVendorID is a method that returns the total bookings by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the total bookings and an error if any.
func (*BookingRepository) GetTotalUsingVendorID(vendorID uint) (int64, error) {
	// count is a placeholder for the count
	var count int64

	// Get the bookings from the database
	err := mysql.Conn.Model(&models.Booking{}).Preload("Order", "status = ?", enums.Success.Label()).Where("vendor_id = ?", vendorID).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting total bookings using vendor id: " + err.Error())

		return 0, err
	}

	return count, nil
}

// GetTotalTodayUsingVendorID is a method that returns the total bookings today by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the total bookings today and an error if any.
func (*BookingRepository) GetTotalTodayUsingVendorID(vendorID uint) (int64, error) {
	// count is a placeholder for the count
	var count int64

	// Get the current date
	today := time.Now().Truncate(24 * time.Hour)

	// Get the bookings from the database
	err := mysql.Conn.Model(&models.Booking{}).Preload("Order", "status = ?", enums.Success.Label()).Where("vendor_id = ? AND date >= ? AND date < ?", vendorID, today, today.Add(24*time.Hour)).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting total bookings today using vendor id: " + err.Error())

		return 0, err
	}

	return count, nil
}

// GetNLatestUsingVendorID is a method that returns the n latest bookings by the given vendor ID.
//
// vendorID: The ID of the vendor.
// n: The number of bookings to return.
//
// Returns the n latest bookings and an error if any.
func (*BookingRepository) GetNLatestUsingVendorID(vendorID uint, n int) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	// Get the bookings from the database
	err := mysql.Conn.Preload("Order", "status = ?", enums.Success.Label()).Where("vendor_id = ?", vendorID).Order("date desc").Limit(n).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		log.Printf("Error getting %d latest bookings using vendor id: %v\n", n, err.Error())

		return nil, err
	}

	return &bookings, nil
}

// CheckUserHasBookCourt is a method that checks if the user has booked the court.
//
// userID: The ID of the user.
// vendorID: The ID of the vendor.
// CheckUserHasBookCourt: The ID of the court.
//
// Returns true if the user has booked the court and an error if any.
func (*BookingRepository) CheckUserHasBookCourt(userID uint, vendorID uint, courtID uint) (bool, error) {
	// count is a placeholder for the count
	var count int64

	// Get the bookings from the database
	err := mysql.Conn.Model(&models.Booking{}).Where("user_id = ? AND vendor_id = ? AND court_id = ?", userID, vendorID, courtID).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking if user has booked the court: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

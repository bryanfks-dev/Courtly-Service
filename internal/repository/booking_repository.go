package repository

import (
	"log"
	"main/core/enums"
	"main/data/models"
	"main/internal/providers/mysql"

	"gorm.io/gorm"
)

// BookingRepository is a struct that defines the BookingRepository
type BookingRepository struct{}

// NewBookingRepository is a function that returns a new BookingRepository
//
// Returns a pointer to the BookingRepository struct
func NewBookingRepository() *BookingRepository {
	return &BookingRepository{}
}

// Create is a method that creates a booking in the database.
//
// tx: The database transaction.
// booking: The booking to create.
//
// Returns an error if any.
func (*BookingRepository) Create(tx *gorm.DB, booking *models.Booking) error {
	// Create the booking in the database
	err := tx.Create(booking).Error

	// Return an error if any
	if err != nil {
		log.Println("Error creating booking: " + err.Error())

		return err
	}

	return nil
}

// CheckUserHasBookCourt is a method that checks if the user has booked the court.
//
// userID: The ID of the user.
// vendorID: The ID of the vendor.
// courtType: The type of the court.
//
// Returns true if the user has booked the court and an error if any.
func (*BookingRepository) CheckUserHasBookCourt(userID uint, vendorID uint, courtType string) (bool, error) {
	// count is a placeholder for the count
	var count int64

	// Get the bookings from the database
	err :=
		mysql.Conn.Model(&models.Booking{}).Joins("JOIN courts ON courts.id = bookings.court_id").
			Joins("JOIN court_types ON court_types.id = courts.court_type_id").
			Where("user_id = ? AND bookings.vendor_id = ? AND court_types.type = ?", userID, vendorID, courtType).
			Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking if user has booked the court: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

// CheckAvailability is a method that checks if the court is available.
//
// courtID: The ID of the court.
// bookDate: The date of the booking.
// bookStartTime: The start time of the booking.
//
// Returns true if the court is available and an error if any.
func (*BookingRepository) CheckAvailability(courtID uint, bookDate string, bookStartTime string) (bool, error) {
	// count is a placeholder for the count
	var count int64

	// Get the bookings from the database
	err :=
		mysql.Conn.Model(&models.Booking{}).Where("court_id = ? AND date = ? AND book_start_time = ?", courtID, bookDate, bookStartTime).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking if a user has booked the court: " + err.Error())

		return false, err
	}

	return count == 0, nil
}

// GetusingVendorIDCourtTypeDate is a method to get bookings using vendor id, court type, and date.
//
// vendorID: the id of the vendor
// courtType: the type of the court
// date: the date of booking
//
// Returns bookings data and error if any
func (*BookingRepository) GetUsingVendorIDCourtTypeDate(vendorID uint, courtType string, date string) (*[]models.Booking, error) {
	// bookings is a placeholder for the bookings
	var bookings []models.Booking

	err :=
		mysql.Conn.Joins("Court").Preload("Court.Vendor").Joins("JOIN orders ON orders.id = bookings.order_id").Where("orders.status = ?", enums.Success.Label()).Where("bookings.vendor_id = ?", vendorID).Where("Court.court_type_id = ?", enums.GetCourtTypeID(courtType)).Where("date = ?", date).Find(&bookings).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting bookings using court id and date: " + err.Error())

		return nil, err
	}

	return &bookings, nil
}

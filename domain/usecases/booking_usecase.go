package usecases

import (
	"main/data/models"
	"main/internal/repository"
)

// BookingUseCase is a struct that defines the BookingUseCase
type BookingUseCase struct {
	BookingRepository *repository.BookingRepository
}

// NewBookingUseCase is a function that returns a new BookingUseCase
//
// b: the booking repository
//
// Returns a pointer to the BookingUseCase struct
func NewBookingUseCase(b *repository.BookingRepository) *BookingUseCase {
	return &BookingUseCase{
		BookingRepository: b,
	}
}

// GetCourtBookings is a use case that gets the court bookings
//
// vendorID: the id of the vendor
// courtType: the court type
// date: The date of the booking
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCourtBookings(vendorID uint, courtType string, date string) (*[]models.Booking, error) {
	return b.BookingRepository.GetUsingVendorIDCourtTypeDate(vendorID, courtType, date)
}

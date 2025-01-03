package usecases

import (
	"main/data/models"
	"main/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

// BookingUseCase is a struct that defines the BookingUseCase
type BookingUseCase struct {
	AuthUseCase       *AuthUseCase
	BookingRepository *repository.BookingRepository
}

// NewBookingUseCase is a function that returns a new BookingUseCase
//
// a: the auth use case
// b: the booking repository
//
// Returns a pointer to the BookingUseCase struct
func NewBookingUseCase(a *AuthUseCase, b *repository.BookingRepository) *BookingUseCase {
	return &BookingUseCase{
		AuthUseCase:       a,
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

// GetCurretnVendorCourtBookings is a use case that gets the current vendor court bookings
// using the token, court type and date
//
// token: the jwt token
// courtType: the court type
// date: The date of the booking
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorCourtBookings(token *jwt.Token, courtType string, date string) (*[]models.Booking, error) {
	// Decode the token
	claims := b.AuthUseCase.DecodeToken(token)

	return b.BookingRepository.GetUsingVendorIDCourtTypeDate(claims.Id, courtType, date)
}

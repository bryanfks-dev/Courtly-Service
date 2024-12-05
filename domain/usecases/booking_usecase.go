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
// b: The BookingRepository
//
// Returns a pointer to the BookingUseCase struct
func NewBookingUseCase(a *AuthUseCase, b *repository.BookingRepository) *BookingUseCase {
	return &BookingUseCase{
		AuthUseCase:       a,
		BookingRepository: b,
	}
}

// GetCurrentUserBookings is a use case that gets the current user bookings
//
// token: The JWT token
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCurrentUserBookings(token *jwt.Token) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	// Get the bookings from the database
	return b.BookingRepository.GetUsingUserID(claims.Id)
}

// GetVendorBookings is a use case that gets the vendor bookings
// by the vendor ID.
//
// token: The JWT token
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorBookings(token *jwt.Token) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.BookingRepository.GetUsingVendorID(claims.Id)
}

// GetCurrentVendorTotalBookings is a use case that gets the current vendor
//
// token: The JWT token
//
// Returns the total bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorTotalBookings(token *jwt.Token) (int64, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.BookingRepository.GetTotalUsingVendorID(claims.Id)
}

// GetCurrentVendorTotalBookingsToday is a use case that gets the current vendor
// total bookings today.
//
// token: The JWT token
//
// Returns the total bookings today and an error if any
func (b *BookingUseCase) GetCurrentVendorTotalBookingsToday(token *jwt.Token) (int64, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.BookingRepository.GetTotalTodayUsingVendorID(claims.Id)
}

// GetCurrentVendorRecentBookings is a use case that gets the current vendor
// recent bookings.
//
// token: The JWT token
//
// Returns the recent bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorRecentBookings(token *jwt.Token) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.BookingRepository.GetNLatestUsingVendorID(claims.Id, 3)
}

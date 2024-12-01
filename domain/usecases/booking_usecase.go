package usecases

import (
	"log"
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

// GetUserBookings is a use case that gets the user bookings
//
// userID: The ID of the user
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetUserBookings(userID uint) (*[]models.Booking, error) {
	// Get the bookings from the database
	bookings, err := b.BookingRepository.GetByUserID(userID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get user bookings: ", err)

		return nil, err
	}

	return bookings, nil
}

// GetCurrentUserBookings is a use case that gets the current user bookings
func (b *BookingUseCase) GetCurrentUserBookings(token *jwt.Token) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	// Get the user bookings
	bookings, err := b.GetUserBookings(uint(claims.Id))

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

// GetVendorBookings is a use case that gets the vendor bookings
//
// vendorID: The ID of the vendor
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetVendorBookings(vendorID uint) (*[]models.Booking, error) {
	// Get the bookings from the database
	bookings, err := b.BookingRepository.GetByVendorID(vendorID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get vendor bookings: ", err)

		return nil, err
	}

	return bookings, nil
}

// GetVendorBookings is a use case that gets the vendor bookings
//
// token: The JWT token
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorBookings(token *jwt.Token) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	// Get the vendor bookings
	bookings, err := b.GetVendorBookings(uint(claims.Id))

	// Return an error if any
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

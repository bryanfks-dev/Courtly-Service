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

// GetUserBookings is a use case that gets the user bookings by the user ID.
//
// userID: The ID of the user
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetUserBookings(userID uint) (*[]models.Booking, error) {
	// Get the bookings from the database
	bookings, err := b.BookingRepository.GetUsingUserID(userID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get user bookings: ", err)

		return nil, err
	}

	return bookings, nil
}

// GetCurrentUserBookings is a use case that gets the current user bookings
//
// token: The JWT token
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCurrentUserBookings(token *jwt.Token) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.GetUserBookings(uint(claims.Id))
}

// GetVendorBookings is a use case that gets the vendor bookings by the vendor ID.
//
// vendorID: The ID of the vendor
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetVendorBookings(vendorID uint) (*[]models.Booking, error) {
	// Get the bookings from the database
	bookings, err := b.BookingRepository.GetUsingVendorID(vendorID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get vendor bookings: ", err)

		return nil, err
	}

	return bookings, nil
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

	return b.GetVendorBookings(uint(claims.Id))
}

// GetVendorTotalBookings is a use case that gets the vendor
// total bookings by the vendor ID.
//
// vendorID: The ID of the vendor
//
// Returns the total bookings and an error if any
func (b *BookingUseCase) GetVendorTotalBookings(vendorID uint) (int, error) {
	// Get the bookings from the database
	total, err := b.BookingRepository.GetTotalUsingVendorID(vendorID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get vendor total bookings: ", err)

		return 0, err
	}

	return int(total), nil
}

// GetCurrentVendorTotalBookings is a use case that gets the current vendor
//
// token: The JWT token
//
// Returns the total bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorTotalBookings(token *jwt.Token) (int, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.GetVendorTotalBookings(uint(claims.Id))
}

// GetVendorTotalBookingsToday is a use case that gets the vendor
// total bookings today by the vendor ID.
//
// vendorID: The ID of the vendor
//
// Returns the total bookings today and an error if any
func (b *BookingUseCase) GetVendorTotalBookingsToday(vendorID uint) (int, error) {
	// Get the bookings from the database
	total, err := b.BookingRepository.GetTotalTodayUsingVendorID(vendorID)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get vendor total bookings today: ", err)

		return 0, err
	}

	return int(total), nil
}

// GetCurrentVendorTotalBookingsToday is a use case that gets the current vendor
// total bookings today.
//
// token: The JWT token
//
// Returns the total bookings today and an error if any
func (b *BookingUseCase) GetCurrentVendorTotalBookingsToday(token *jwt.Token) (int, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.GetVendorTotalBookingsToday(uint(claims.Id))
}

// GetVendorRecentBookings is a use case that gets the vendor recent bookings
// by the vendor ID.
//
// vendorID: The ID of the vendor
//
// Returns the recent bookings and an error if any
func (b *BookingUseCase) GetVendorRecentBookings(vendorID uint) (*[]models.Booking, error) {
	// Get the bookings from the database
	bookings, err := b.BookingRepository.GetNLatestUsingVendorID(vendorID, 5)

	// Return an error if any
	if err != nil {
		log.Println("Failed to get vendor recent bookings: ", err)

		return nil, err
	}

	return bookings, nil
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

	return b.GetVendorRecentBookings(uint(claims.Id))
}

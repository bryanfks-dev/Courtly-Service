package usecases

import (
	"context"
	"main/core/types"
	"main/data/models"
	"main/internal/repository"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

// BookingUseCase is a struct that defines the BookingUseCase
type BookingUseCase struct {
	AuthUseCase       *AuthUseCase
	BookingRepository *repository.BookingRepository
	CourtRepository   *repository.CourtRepository
}

// NewBookingUseCase is a function that returns a new BookingUseCase
//
// a: The AuthUseCase
// b: The BookingRepository
// c: The CourtRepository
//
// Returns a pointer to the BookingUseCase struct
func NewBookingUseCase(a *AuthUseCase, b *repository.BookingRepository, c *repository.CourtRepository) *BookingUseCase {
	return &BookingUseCase{
		AuthUseCase:       a,
		BookingRepository: b,
		CourtRepository:   c,
	}
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

// GetCurrentVendorOrdersStats is a use case that gets the current vendor orders
// statistics from the database.
//
// token: The JWT token
//
// Returns the map of order stat and an error if any
func (b *BookingUseCase) GetCurrentVendorOrdersStats(token *jwt.Token) (*types.OrdersStatsMap, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	// Placeholder for the bookings and error
	stats := make(types.OrdersStatsMap)

	// Get the total bookings
	var err error

	// Create a context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	// Defer the cancel function
	defer cancel()

	// Create a wait group
	var (
		wg sync.WaitGroup
	)

	// Add a wait group
	wg.Add(1)

	// Get the total bookings
	go func() {
		// Defer done
		defer wg.Done()

		// Get the total bookings
		totalOrders, e := b.BookingRepository.GetTotalUsingVendorID(claims.Id)

		// Return an error if any
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the total orders to the stats
		stats["total_orders"] = totalOrders
	}()

	// Add a wait group
	wg.Add(1)

	// Get the total bookings today
	go func() {
		// Defer done
		defer wg.Done()

		// Get the total bookings today
		totalOrdersToday, e := b.BookingRepository.GetTotalTodayUsingVendorID(claims.Id)

		// Return an error if any
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the total orders today to the stats
		stats["total_orders_today"] = totalOrdersToday
	}()

	// Add a wait group
	wg.Add(1)

	// Get the recent bookings
	go func() {
		// Defer done
		defer wg.Done()

		// Get the recent bookings
		recentBooking, e := b.BookingRepository.GetNLatestUsingVendorID(claims.Id, 3)

		// Return an error if any
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the recent bookings to the stats
		stats["recent_orders"] = recentBooking
	}()

	// Wait for the wait group to finish
	wg.Wait()

	// Wait for the wait group to finish
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetCurrentVendorBookingsUsingCourtType is a use case that gets the current vendor bookings
// by the court type.
//
// token: The JWT token
//
// Returns the bookings and an error if any
func (b *BookingUseCase) GetCurrentVendorBookingsUsingCourtType(token *jwt.Token, courtType string) (*[]models.Booking, error) {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	return b.BookingRepository.GetUsingVendorIDCourtType(claims.Id, courtType)
}

package usecases

import (
	"context"
	"main/core/constants"
	"main/core/shared"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/providers/mysql"
	"main/internal/repository"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// BookingUseCase is a struct that defines the BookingUseCase
type BookingUseCase struct {
	AuthUseCase       *AuthUseCase
	BookingRepository *repository.BookingRepository
	CourtRepository   *repository.CourtRepository
	OrderRepository   *repository.OrderRepository
}

// NewBookingUseCase is a function that returns a new BookingUseCase
//
// a: The AuthUseCase
// b: The BookingRepository
// c: The CourtRepository
//
// Returns a pointer to the BookingUseCase struct
func NewBookingUseCase(a *AuthUseCase, b *repository.BookingRepository, c *repository.CourtRepository, o *repository.OrderRepository) *BookingUseCase {
	return &BookingUseCase{
		AuthUseCase:       a,
		BookingRepository: b,
		CourtRepository:   c,
		OrderRepository:   o,
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

// ValidateCreateBooking is a use case that validates the create booking data
//
// data: The booking data
//
// Returns an error message if any
func (b *BookingUseCase) ValidateCreateBooking(data dto.CreateBookingDTO) string {
	// Check if the vendor ID is empty
	if data.VendorID == 0 {
		return "Vendor ID is required"
	}

	// Check if the date is empty
	if data.Date == "" {
		return "Date is required"
	}

	// Check if the bookings is empty
	if data.Bookings == nil {
		return "Bookings is required"
	}

	// Loop through the bookings
	for _, booking := range *data.Bookings {
		// Check if the court ID is empty
		if booking.CourtID == 0 {
			return "Court ID is required"
		}

		// Check if the book time is empty
		if booking.BookTime == nil {
			return "Book time is required"
		}

		// Loop through the book time
		for _, bookTime := range booking.BookTime {
			// Check if the book time is empty
			if bookTime == "" {
				return "Book time is required"
			}
		}
	}

	return ""
}

// CreateBooking is a use case that creates a booking
// by the user.
//
// token: The JWT token
// data: The booking data
//
// Returns an error if any
func (b *BookingUseCase) CreateBooking(token *jwt.Token, data dto.CreateBookingDTO) *entities.ProcessError {
	// Get the token claims
	claims := b.AuthUseCase.DecodeToken(token)

	// Parse the date
	parsedDate, err := time.Parse("2006-01-02", data.Date)

	// Return an error if any
	if err != nil {
		return &entities.ProcessError{
			ClientError: true,
			Message:     "Invalid date format",
		}
	}

	// Begin a transaction
	tx := mysql.Conn.Begin()

	// Defer the rollback
	defer func() {
		// Recover from panic
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Return an error if any
	if tx.Error != nil {
		return &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to begin transaction",
		}
	}

	// Create a context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	// Defer the cancel function
	defer cancel()

	// Create a wait group
	var (
		wg         sync.WaitGroup
		processErr *entities.ProcessError
	)

	// Get the court
	court, err := b.CourtRepository.GetUsingID((*data.Bookings)[0].CourtID)

	// Return an error if any
	if err != nil {
		return &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to get court",
		}
	}

	// Create Order for bookings
	order := models.Order{
		Price:           court.Price * float64(len(*data.Bookings)),
		PaymentMethodID: nil,
		AppFee:          constants.APP_FEE_PRICE,
		Status:          "Pending",
	}

	// Create the order
	err = b.OrderRepository.Create(tx, &order)

	// Return an error if any
	if err != nil {
		// Rollback the transaction
		tx.Rollback()

		return &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to create order",
		}
	}

	// Loop through the bookings
	for _, booking := range *data.Bookings {
		// Add a wait group
		wg.Add(1)

		go func(booking dto.CreateBookingDTOInner) {
			defer wg.Done()

			// Create a wait group
			var innerWg sync.WaitGroup

			// Loop through the booking times
			for _, bookTime := range booking.BookTime {
				// Add a wait group
				innerWg.Add(1)

				go func(bookTime string) {
					defer innerWg.Done()

					parsedTime, e := time.Parse("15:04", bookTime)

					// Return an error if any
					if e != nil {
						// Set the process error
						processErr = &entities.ProcessError{
							ClientError: true,
							Message:     "Invalid time format",
						}

						// Rollback the transaction
						tx.Rollback()

						cancel()

						return
					}

					// Create a new booking
					book := models.Booking{
						UserID:        claims.Id,
						VendorID:      data.VendorID,
						OrderID:       order.ID,
						CourtID:       booking.CourtID,
						Date:          shared.DateOnly{Time: parsedDate},
						BookStartTime: shared.TimeOnly{Time: parsedTime},
						BookEndTime:   shared.TimeOnly{Time: parsedTime.Add(time.Hour)},
					}

					// Return an error if any
					e = func() error {
						mu := sync.Mutex{}

						// Lock the database
						mu.Lock()
						defer mu.Unlock()

						return b.BookingRepository.Create(tx, &book)
					}()

					// Return an error if any
					if e != nil {
						// Set the process error
						processErr = &entities.ProcessError{
							ClientError: false,
							Message:     "Failed to create booking",
						}

						// Rollback the transaction
						tx.Rollback()

						cancel()

						return
					}
				}(bookTime)
			}

			// Wait all create booking
			innerWg.Wait()
		}(booking)
	}

	// Wait all court iteration
	wg.Wait()

	// Return an error if any
	if processErr != nil {
		return processErr
	}

	// Return an error if any
	if err := tx.Commit().Error; err != nil {
		return &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to commit transaction",
		}
	}

	return nil
}

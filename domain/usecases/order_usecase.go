package usecases

import (
	"context"
	"main/core/constants"
	"main/core/shared"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/providers/midtrans"
	"main/internal/providers/mysql"
	"main/internal/repository"
	"main/pkg/utils"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// OrderUseCase is a struct that defines the OrderUseCase
type OrderUseCase struct {
	AuthUseCase       *AuthUseCase
	OrderRepository   *repository.OrderRepository
	BookingRepository *repository.BookingRepository
	CourtRepository   *repository.CourtRepository
}

// NewOrderUseCase is a function that returns a new OrderUseCase
//
// a: The AuthUseCase
// o: The OrderRepository
// b: The BookingRepository
// c: The CourtRepository
//
// Returns a pointer to the OrderUseCase struct
func NewOrderUseCase(a *AuthUseCase, o *repository.OrderRepository, b *repository.BookingRepository, c *repository.CourtRepository) *OrderUseCase {
	return &OrderUseCase{
		AuthUseCase:       a,
		OrderRepository:   o,
		BookingRepository: b,
		CourtRepository:   c,
	}
}

// GetCurrentUserOrders is a method that gets the current user order from the database.
//
// token: The JWT token.
// courtType: The court type.
//
// Returns the orders and an error if any.
func (o *OrderUseCase) GetCurrentUserOrders(token *jwt.Token, courtType *string) (*[]models.Order, error) {
	// Get the user ID from the JWT
	claims := o.AuthUseCase.DecodeToken(token)

	// Check if the court type is not empty
	if courtType != nil && utils.IsBlank(*courtType) {
		// Get the orders using the user ID
		return o.OrderRepository.GetUsingUserID(claims.Id)
	}

	// Get the orders using the user ID
	return o.OrderRepository.GetUsingUserIDCourtType(claims.Id, *courtType)
}

// GetCurrentVendorOrders is a method that gets the current vendor order from the database.
//
// token: The JWT token.
// courtType: The court type.
//
// Returns the orders and an error if any.
func (o *OrderUseCase) GetCurrentVendorOrders(token *jwt.Token, courtType *string) (*[]models.Order, error) {
	// Get the user ID from the JWT
	claims := o.AuthUseCase.DecodeToken(token)

	// Check if the court type is not empty
	if courtType != nil && utils.IsBlank(*courtType) {
		// Get the orders using the user ID
		return o.OrderRepository.GetUsingVendorID(claims.Id)
	}

	// Get the orders using the user ID
	return o.OrderRepository.GetUsingVendorIDCourtType(claims.Id, *courtType)
}

// GetCurrentUserOrderDetail is a method that gets the current user order detail from the database.
//
// token: The JWT token.
// orderID: The order ID.
//
// Returns the order and an error if any.
func (o *OrderUseCase) GetCurrentUserOrderDetail(token *jwt.Token, orderID uint) (*models.Order, *entities.ProcessError) {
	// Get the order using the order ID
	order, err := o.OrderRepository.GetUsingID(orderID)

	// Return an error if any
	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "Order not found",
		}
	}

	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to get user order",
		}
	}

	// Get the user ID from the JWT
	claims := o.AuthUseCase.DecodeToken(token)

	// Return an error if the order is not belongs to the user
	if order.Bookings[0].UserID != claims.Id {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "This order is not belongs to this user",
		}
	}

	return order, nil
}

// GetCurrentVendorOrderDetail is a method that gets the current vendor order detail from the database.
//
// token: The JWT token.
// orderID: The order ID.
//
// Returns the order and an error if any.
func (o *OrderUseCase) GetCurrentVendorOrderDetail(token *jwt.Token, orderID uint) (*models.Order, *entities.ProcessError) {
	// Get the order using the order ID
	order, err := o.OrderRepository.GetUsingID(orderID)

	// Return an error if any
	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "Order not found",
		}
	}

	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to get user order",
		}
	}

	// Get the user ID from the JWT
	claims := o.AuthUseCase.DecodeToken(token)

	// Return an error if the order is not belongs to the user
	if order.Bookings[0].VendorID != claims.Id {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "This order is not belongs to this vendor",
		}
	}

	return order, nil
}

// ValidateCreateOrder is a use case that validates the create order data
//
// data: The order data
//
// Returns an error message if any
func (o *OrderUseCase) ValidateCreateOrder(data dto.CreateOrderDTO) string {
	// Check if the vendor ID is empty
	if data.VendorID == 0 {
		return "Vendor ID is required"
	}

	// Check if the date is empty
	if utils.IsBlank(data.Date) {
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
			if utils.IsBlank(bookTime) {
				return "Book time is required"
			}

			// Check if the book time is valid
			available, err := o.BookingRepository.CheckAvailability(booking.CourtID, data.Date, bookTime)

			// Return an error if any
			if err != nil {
				return "Failed to check availability"
			}

			// Return an error if the court is not available
			if !available {
				return "Court is not available at this time"
			}
		}
	}

	return ""
}

// CreateOrder is a use case that creates an order
// by the user.
//
// token: The JWT token
// data: The order data
//
// Returns the payment token and error if any
func (o *OrderUseCase) CreateOrder(token *jwt.Token, data dto.CreateOrderDTO) (*string, *entities.ProcessError) {
	// Get the token claims
	claims := o.AuthUseCase.DecodeToken(token)

	// Parse the date
	parsedDate, err := time.Parse("2006-01-02", data.Date)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
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
		return nil, &entities.ProcessError{
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
	court, err := o.CourtRepository.GetUsingID((*data.Bookings)[0].CourtID)

	// Return an error if any
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to get court",
		}
	}

	// Get the total bookings
	totalBookings := 0

	for i := 0; i < len(*data.Bookings); i++ {
		totalBookings += len((*data.Bookings)[i].BookTime)
	}

	// Create Order for bookings
	order := models.Order{
		Price:  court.Price * float64(totalBookings),
		AppFee: constants.APP_FEE_PRICE,
		Status: "Pending",
	}

	// Create the order
	err = o.OrderRepository.Create(tx, &order)

	// Return an error if any
	if err != nil {
		// Rollback the transaction
		tx.Rollback()

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to create order",
		}
	}

	// Create a payment token
	paymentToken, err := midtrans.CreateToken(order.ID, int64(order.Price))

	// Return an error if any
	if err != nil {
		// Rollback the transaction
		tx.Rollback()

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to create transaction",
		}
	}

	// Update the order payment token
	err = o.OrderRepository.UpdatePaymentTokenUsingID(tx, *paymentToken, order.ID)

	// Return an error if any
	if err != nil {
		// Rollback the transaction
		tx.Rollback()

		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to update payment token",
		}
	}

	// Loop through the bookings
	for _, booking := range *data.Bookings {
		// Add a wait group
		wg.Add(1)

		go func(booking dto.CreateOrderDTOInner) {
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

						return o.BookingRepository.Create(tx, &book)
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
		return nil, processErr
	}

	// Return an error if any
	if err := tx.Commit().Error; err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to commit transaction",
		}
	}

	return paymentToken, nil
}

// GetCurrentVendorOrdersStats is a use case that gets the current vendor orders
// statistics from the database.
//
// token: The JWT token
//
// Returns the map of order stat and an error if any
func (o *OrderUseCase) GetCurrentVendorOrdersStats(token *jwt.Token) (*types.OrdersStatsMap, error) {
	// Get the token claims
	claims := o.AuthUseCase.DecodeToken(token)

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
		totalOrders, e := o.OrderRepository.GetTotalUsingVendorID(claims.Id)

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
		totalOrdersToday, e := o.OrderRepository.GetTotalTodayUsingVendorID(claims.Id)

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
		recentBooking, e :=
			o.OrderRepository.GetNLatestUsingVendorID(claims.Id, constants.LATEST_ORDER_LIMIT)

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

package controllers

import (
	"log"
	"main/core/enums"
	"main/data/models"
	"main/domain/usecases"
	"main/internal/dto"
	"main/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BookingController is a struct that defines the BookingController
type BookingController struct {
	BookingUseCase *usecases.BookingUseCase
}

// NewBookingController is a function that returns a new BookingController
//
// b: The BookingUseCase
//
// Returns a pointer to the BookingController struct
func NewBookingController(b *usecases.BookingUseCase) *BookingController {
	return &BookingController{
		BookingUseCase: b,
	}
}

// GetCurrentUserBookings is a controller that gets the current user bookings
// from the database.
// Endpoint: GET /users/me/bookings
//
// c: The echo context.
//
// Returns an error if any.
func (b *BookingController) GetCurrentUserBookings(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the court type from the query parameter
	courtTypeParam := c.QueryParam("type")

	// Placeholder for the bookings and error
	var (
		bookings *[]models.Booking
		err      error
	)

	// Check if the court type is not empty
	if !utils.IsBlank(courtTypeParam) && !enums.InCourtType(courtTypeParam) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get the current user bookings
	// Check if the court type is not empty
	if utils.IsBlank(courtTypeParam) {
		bookings, err = b.BookingUseCase.GetCurrentUserBookings(cc.Token)
	} else {
		bookings, err = b.BookingUseCase.GetCurrentUserBookingsUsingCourtType(cc.Token, courtTypeParam)
	}

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get user bookings",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User bookings retrieved successfully",
		Data:    dto.CurrentUserBookingsResponseDTO{}.FromModels(bookings),
	})
}

// GetCurrentUserBooking is a controller that gets the current user booking
// from the database.
// Endpoint: GET /vendors/me/orders
//
// c: The echo context.
//
// Returns an error if any.
func (b *BookingController) GetCurrentVendorOrders(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the court type from the query parameter
	courtTypeParam := c.QueryParam("type")

	// Placeholder for the bookings and error
	var (
		bookings *[]models.Booking
		err      error
	)

	// Check if the court type is not empty
	if !utils.IsBlank(courtTypeParam) && !enums.InCourtType(courtTypeParam) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get the current vendor bookings
	// Check if the court type is not empty
	if utils.IsBlank(courtTypeParam) {
		bookings, err = b.BookingUseCase.GetCurrentVendorBookings(cc.Token)
	} else {
		bookings, err = b.BookingUseCase.GetCurrentVendorBookingsUsingCourtType(cc.Token, courtTypeParam)
	}

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get vendor orders",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Vendor orders retrieved successfully",
		Data:    dto.CurrentVendorOrdersResponseDTO{}.FromModels(bookings),
	})
}

// GetCurrentVendorOrdersStats is a controller that gets the current vendor orders
// statistics from the database.
// Endpoint: GET /vendors/me/orders/stats
//
// c: The echo context.
//
// Returns an error if any.
func (b *BookingController) GetCurrentVendorOrdersStats(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get current vendor orders stats
	stats, err := b.BookingUseCase.GetCurrentVendorOrdersStats(cc.Token)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get vendor orders stats",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Vendor orders stats retrieved successfully",
		Data:    dto.CurrentVendorOrdersStatsResponseDTO{}.FromMap(stats),
	})
}

// CreateBooking is a controller that creates a new booking.
// Endpoint: POST /users/me/bookings
//
// c: The echo context.
//
// Returns an error if any.
func (b *BookingController) CreateBooking(c echo.Context) error {
	// Create a new CreateBookingDTO object
	data := new(dto.CreateBookingDTO)

	// Bind the request body to the CreateBookingDTO object
	if err := c.Bind(data); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Validate the CreateBookingDTO object
	errorMsg := b.BookingUseCase.ValidateCreateBooking(*data)

	// Return an error if any
	if !utils.IsBlank(errorMsg) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: errorMsg,
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Create a new booking
	err := b.BookingUseCase.CreateBooking(cc.Token, *data)

	// Return an error if any
	if err != nil {
		// Check if the error is a client error
		if err.ClientError {
			return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
				Success: false,
				Message: err.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: err.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Booking created successfully",
		Data:    nil,
	})
}

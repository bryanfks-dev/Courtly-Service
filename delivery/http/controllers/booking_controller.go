package controllers

import (
	"main/domain/usecases"
	"main/internal/dto"
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
// GET /users/me/bookings
//
// c: The echo context.
//
// Returns an error if any.
func (b *BookingController) GetCurrentUserBookings(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current user
	bookings, err := b.BookingUseCase.GetCurrentUserBookings(cc.Token)

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
// GET /vendors/me/orders
//
// c: The echo context.
//
// Returns an error if any.
func (b *BookingController) GetCurrentVendorOrders(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor bookings
	bookings, err := b.BookingUseCase.GetCurrentVendorBookings(cc.Token)

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

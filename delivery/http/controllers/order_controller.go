package controllers

import (
	"log"
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"main/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// OrderController is a struct that defines the OrderController
type OrderController struct {
	OrderUseCase  *usecases.OrderUseCase
	ReviewUseCase *usecases.ReviewUseCase
}

// NewOrderController is a function that returns a new OrderController
//
// o: The OrderUseCase
// r: The ReviewUseCase
//
// Returns a pointer to the OrderController struct
func NewOrderController(o *usecases.OrderUseCase, r *usecases.ReviewUseCase) *OrderController {
	return &OrderController{
		OrderUseCase:  o,
		ReviewUseCase: r,
	}
}

// GetCurrentUserBooking is a controller that gets the current user booking
// from the database.
// Endpoint: GET /vendors/me/orders
//
// c: The echo context.
//
// Returns an error if any.
func (o *OrderController) GetCurrentVendorOrders(c echo.Context) error {
	// Get the court type from the query parameter
	courtTypeParam := c.QueryParam("type")

	// Check if the court type is not empty
	if !utils.IsBlank(courtTypeParam) && !enums.InCourtType(courtTypeParam) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor orders
	orders, err :=
		o.OrderUseCase.GetCurrentVendorOrders(cc.Token, &courtTypeParam)

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
		Data:    dto.CurrentVendorOrdersResponseDTO{}.FromModels(orders),
	})
}

// GetCurrentVendorOrdersStats is a controller that gets the current vendor orders
// statistics from the database.
// Endpoint: GET /vendors/me/orders/stats
//
// c: The echo context.
//
// Returns an error if any.
func (o *OrderController) GetCurrentVendorOrdersStats(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get current vendor orders stats
	stats, err := o.OrderUseCase.GetCurrentVendorOrdersStats(cc.Token)

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

// CreateOrder is a controller that creates a new booking.
// Endpoint: POST /users/me/orders
//
// c: The echo context.
//
// Returns an error if any.
func (o *OrderController) CreateOrder(c echo.Context) error {
	// Create a new CreateBookingDTO object
	data := new(dto.CreateOrderDTO)

	// Bind the request body to the CreateBookingDTO object
	if err := c.Bind(data); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Validate the CreateBookingDTO object
	errorMsg := o.OrderUseCase.ValidateCreateOrder(*data)

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
	paymentToken, err := o.OrderUseCase.CreateOrder(cc.Token, *data)

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
		Message: "Order created successfully",
		Data: dto.CreateOrderResponseDTO{
			PaymentToken: *paymentToken,
		},
	})
}

// GetCurrentUserOrders is a controller that gets the current user orders
// from the database.
// Endpoint: GET /users/me/orders
//
// c: The echo context.
//
// Returns an error if any.
func (o *OrderController) GetCurrentUserOrders(c echo.Context) error {
	// Get the court type from the query parameter
	courtType := c.QueryParam("type")

	// Check if the court type is not empty
	if !utils.IsBlank(courtType) && !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current user orders
	orders, err := o.OrderUseCase.GetCurrentUserOrders(cc.Token, &courtType)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get user orders",
			Data:    nil,
		})
	}

	// Check if the orders is empty
	if len(*orders) == 0 {
		return c.JSON(http.StatusOK, dto.ResponseDTO{
			Success: true,
			Message: "User orders retrieved successfully",
			Data:    dto.CurrentUserOrdersResponseDTO{}.FromModels(orders),
		})
	}

	// Create order dtos
	// Create a slice of order DTOs
	dtos := []dto.CurrentUserOrderDTO{}

	// Convert the orders to order DTOs
	for _, order := range *orders {
		// Check if user has reviewed for court type
		reviewed, err := o.ReviewUseCase.CheckCurrentUserHasReviewedUsingVendorIDCourtType(cc.Token, order.Bookings[0].VendorID, order.Bookings[0].Court.CourtType.Type)

		// Return an error if any
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get user orders",
				Data:    nil,
			})
		}

		// Append the order DTO to the slice
		dtos = append(dtos, *dto.CurrentUserOrderDTO{}.FromModel(&order, &reviewed))
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User orders retrieved successfully",
		Data: dto.CurrentUserOrdersResponseDTO{
			Orders: &dtos,
		},
	})
}

// GetCurrentUserOrderDetail is a controller that gets the current user order detail
// from the database.
// Endpoint: GET /users/me/orders/:id
//
// c: The echo context.
//
// Returns an error if any.
func (o *OrderController) GetCurrentUserOrderDetail(c echo.Context) error {
	// Get the order ID from the path parameter
	id := c.Param("id")

	// Check if the id is not empty
	if utils.IsBlank(id) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Order id is required",
			Data:    nil,
		})
	}

	// Convert the order ID to uint
	orderID, err := strconv.Atoi(id)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid order ID",
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current user order detail
	order, processErr :=
		o.OrderUseCase.GetCurrentUserOrderDetail(cc.Token, uint(orderID))

	// Return an error if any
	if processErr != nil {
		if processErr.ClientError {
			return c.JSON(http.StatusNotFound, dto.ResponseDTO{
				Success: false,
				Message: processErr.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: processErr.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User order detail retrieved successfully",
		Data: dto.CurrentUserOrderDetailResponseDTO{
			OrderDetail: dto.CurrentUserOrderDetailDTO{}.FromModel(order),
		},
	})
}

// GetCurrentVendorOrderDetail is a controller that gets the current vendor order detail
// from the database.
// Endpoint: GET /vendors/me/orders/:id
//
// c: The echo context.
//
// Returns an error if any.
func (o *OrderController) GetCurrentVendorOrderDetail(c echo.Context) error {
	// Get the order ID from the path parameter
	id := c.Param("id")

	// Check if the id is not empty
	if utils.IsBlank(id) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Order id is required",
			Data:    nil,
		})
	}

	// Convert the order ID to uint
	orderID, err := strconv.Atoi(id)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid order ID",
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current user order detail
	order, processErr :=
		o.OrderUseCase.GetCurrentVendorOrderDetail(cc.Token, uint(orderID))

	// Return an error if any
	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusNotFound, dto.ResponseDTO{
				Success: false,
				Message: processErr.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: processErr.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User order detail retrieved successfully",
		Data: dto.CurrentVendorOrderDetailResponseDTO{
			OrderDetail: dto.CurrentVendorOrderDetailDTO{}.FromModel(order),
		},
	})
}

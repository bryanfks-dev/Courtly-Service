package controllers

import (
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
	OrderUseCase *usecases.OrderUseCase
}

// NewOrderController is a function that returns a new OrderController
//
// o: The OrderUseCase
//
// Returns a pointer to the OrderController struct
func NewOrderController(o *usecases.OrderUseCase) *OrderController {
	return &OrderController{
		OrderUseCase: o,
	}
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

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User orders retrieved successfully",
		Data:    dto.CurrentUserOrdersDTO{}.FromModels(orders),
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
			Message: "Invalid order ID",
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
			OrderDetail: dto.OrderDetailDTO{}.FromModel(order),
		},
	})
}

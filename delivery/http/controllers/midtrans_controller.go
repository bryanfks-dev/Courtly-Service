package controllers

import (
	"main/internal/dto"
	"main/internal/providers/midtrans"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MidtransController is a struct that defines the MidtransController
type MidtransController struct {}

// NewMidtransController is a function that returns a new MidtransController
//
// Returns a pointer to the MidtransController struct
func NewMidtransController() *MidtransController {
	return &MidtransController{}
}

// PaymentCallback is a controller that handles the payment callback from Midtrans
// Endpoint: POST /midtrans/payment-callback
//
// c: The echo context.
//
// Returns an error if any.
func (m *MidtransController) PaymentCallback(c echo.Context) error {
	// Get the payload from the request
	var notificationPayload map[string]any

	// Bind the payload to the notificationPayload
	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid payload",
		})
	}

	// Verify the payment
	err := midtrans.VerifyPayment(notificationPayload)

	// Check if there is an error
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Error verifying payment",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Payment verified",
		Data: nil,
	})
}

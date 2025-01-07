package controllers

import (
	"main/core/constants"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// FeesController is a struct that defines the methods
// of the FeesController
type FeesController struct{}

// NewFeesController is a function that returns a new FeesController
//
// Returns a new FeesController
func NewFeesController() *FeesController {
	return &FeesController{}
}

// GetFees is a function that handles the request to get the fees
// of the application.
//
// e: The echo context
//
// Returns an error response if there is an error, otherwise a success response.
func (f *FeesController) GetFees(c echo.Context) error {
	return c.JSON(http.StatusOK, &dto.ResponseDTO{
		Success: true,
		Message: "Fees retrieved successfully",
		Data: &dto.FeesResponseDTO{
			AppFee: constants.APP_FEE_PRICE,
		},
	})
}

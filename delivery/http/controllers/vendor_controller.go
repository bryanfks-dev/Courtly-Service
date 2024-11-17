package controllers

import (
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VendorController is a struct that defines the VendorController
// and its usecases.
type VendorController struct {
	VendorUseCase *usecases.VendorUseCase
}

// NewVendorController is a factory function that returns a new instance of the VendorController.
//
// v: The vendor use case.
//
// Returns a new instance of the VendorController.
func NewVendorController(v *usecases.VendorUseCase) *VendorController {
	return &VendorController{
		VendorUseCase: v,
	}
}

// GetCurrentVendor is a handler function that returns the current vendor.
// Endpoint: GET api/v1/vendors/me
//
// c: The echo context.
//
// Returns an error if any.
func (v *VendorController) GetCurrentVendor(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor
	vendor, err := v.VendorUseCase.GetCurrentVendor(cc.Token)

	// Return an error if any
	if err != nil {
		// Check if the error is a client error
		if err.ClientError {
			return c.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Message: err.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Vendor retrieved successfully",
		Data: dto.CurrentVendorResponseData{
			Vendor: dto.CurrentVendor{}.FromModel(vendor),
		},
	})
}

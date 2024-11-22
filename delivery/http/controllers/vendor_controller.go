package controllers

import (
	"log"
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
// Endpoint: GET /vendors/me
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
		Message: "Vendor retrieved successfully",
		Data: dto.CurrentVendorResponseDTO{
			Vendor: dto.CurrentVendorDTO{}.FromModel(vendor),
		},
	})
}

// UpdateVendorPassword is a handler function that updates the password of the current vendor.
// Endpoint: PACTH /vendors/me/password
//
// c: The echo context.
//
// Returns an error if any.
func (v *VendorController) UpdateVendorPassword(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Bind the form data
	form := new(dto.ChangePasswordForm)

	// Return an error if the form data is invalid
	if err := c.Bind(form); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Validate the form data
	if err := v.VendorUseCase.ValidateChangePasswordForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Update the password
	vendor, err := v.VendorUseCase.ProcessChangePassword(cc.Token, form)

	// Return an error if any
	if err != nil {
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
		Message: "Password updated successfully",
		Data: dto.CurrentVendorResponseDTO{
			Vendor: dto.CurrentVendorDTO{}.FromModel(vendor),
		},
	})
}

package controllers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VerifyPasswordController is a struct to hold the VerifyPassword function
type VerifyPasswordController struct {
	VerifyPasswordUseCase *usecases.VerifyPasswordUseCase
}

// NewVerifyPasswordController is a factory function that returns a new instance of VerifyPasswordController
//
// v: The VerifyPasswordUseCase
//
// Returns a new instance of VerifyPasswordController
func NewVerifyPasswordController(v *usecases.VerifyPasswordUseCase) *VerifyPasswordController {
	return &VerifyPasswordController{
		VerifyPasswordUseCase: v,
	}
}

// UserVerifyPassword is a controller to handle the request to verify the password of the user
// Endpoint: POST /auth/user/verify-password
//
// c: Context of the HTTP request
//
// Returns an error if any
func (v *VerifyPasswordController) UserVerifyPassword(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Bind the form dto
	form := new(dto.VerifyPasswordFormDTO)

	// Return an error if the form data is invalid
	if err := c.Bind(form); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Validate the form
	if err := v.VerifyPasswordUseCase.ValidateForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Process the form
	user, err := v.VerifyPasswordUseCase.ProcessUser(form, cc.Token)

	// Check if there is an error
	if err != nil {
		// Check if the error is a client error
		if err.ClientError {
			return c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
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
		Message: "Password verified successfully",
		Data: dto.CurrentUserResponseDTO{
			User: dto.CurrentUserDTO{}.FromModel(user),
		},
	})
}

// VendorVerifyPassword is a controller to handle the request to verify the password of the vendor
// Endpoint: POST /auth/vendor/verify-password
//
// c: Context of the HTTP request
//
// Returns an error if any
func (v *VerifyPasswordController) VendorVerifyPassword(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Bind the form dto
	form := new(dto.VerifyPasswordFormDTO)

	// Return an error if the form data is invalid
	if err := c.Bind(form); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Validate the form
	if err := v.VerifyPasswordUseCase.ValidateForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Process the form
	vendor, err := v.VerifyPasswordUseCase.ProcessVendor(form, cc.Token)

	// Check if there is an error
	if err != nil {
		// Check if the error is a client error
		if err.ClientError {
			return c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
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
		Message: "Password verified successfully",
		Data: dto.CurrentVendorResponseDTO{
			Vendor: dto.CurrentVendorDTO{}.FromModel(vendor),
		},
	})
}

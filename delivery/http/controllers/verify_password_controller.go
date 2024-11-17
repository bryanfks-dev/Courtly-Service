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

	// Bind the form
	form := new(dto.VerifyPasswordForm)

	// Return an error if the form data is invalid
	if err := c.Bind(form); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Validate the form
	if err := v.VerifyPasswordUseCase.ValidateForm(form); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Process the form
	user, err := v.VerifyPasswordUseCase.Process(form, cc.Token)

	// Check if there is an error
	if err != nil {
		return c.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Password verified successfully",
		Data: dto.CurrentUserResponseData{
			User: dto.CurrentUser{}.FromModel(user),
		},
	})
}

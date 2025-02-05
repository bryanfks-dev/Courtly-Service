package controllers

import (
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LogoutController is a struct that defines the logout controller.
type LogoutController struct {
	LogoutUseCase *usecases.LogoutUseCase
}

// NewLogoutController is a factory function that returns a new instance of the LogoutController.
//
// l: The logout use case.
//
// Returns a new instance of the LogoutController.
func NewLogoutController(l *usecases.LogoutUseCase) *LogoutController {
	return &LogoutController{LogoutUseCase: l}
}

// UserLogout is a handler that logs out a user
// by blacklisting the token used to authenticate the user.
// Endpoint: POST /auth/user/logout
//
// c: echo.Context
//
// Returns an error response if there is an error, otherwise a success response.
func (l *LogoutController) UserLogout(c echo.Context) error {
	cc := c.(*dto.CustomContext)

	// Blacklist token
	err := l.LogoutUseCase.BlacklistToken(cc.Token)

	// Check if there was an error blacklisting the token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Could not blacklist token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Successfully logged out",
		Data:    nil,
	})
}

// VendorLogout is a handler that logs out a vendor
// by blacklisting the token used to authenticate the vendor.
// Endpoint: POST /auth/vendor/logout
//
// c: echo.Context
//
// Returns an error response if there is an error, otherwise a success response.
func (l *LogoutController) VendorLogout(c echo.Context) error {
	cc := c.(*dto.CustomContext)

	// Blacklist token
	err := l.LogoutUseCase.BlacklistToken(cc.Token)

	// Check if there was an error blacklisting the token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Could not blacklist token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Successfully logged out",
		Data:    nil,
	})
}

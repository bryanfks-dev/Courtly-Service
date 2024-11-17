package controllers

import (
	"log"
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginController is a struct that defines the login controller.
type LoginController struct {
	loginUseCase *usecases.LoginUseCase
	authUseCase  *usecases.AuthUseCase
}

// NewLoginController is a factory function that returns a new instance of the LoginController.
//
// l: The login use case.
// a: The auth use case.
//
// Returns a new instance of the LoginController.
func NewLoginController(l *usecases.LoginUseCase, a *usecases.AuthUseCase) *LoginController {
	return &LoginController{loginUseCase: l, authUseCase: a}

}

// UserLogin is a function that handles the user login request.
// Endpoint: POST /api/v1/auth/user/login
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func (l *LoginController) UserLogin(c echo.Context) error {
	// Create a new UserLoginForm object
	form := new(dto.UserLoginForm)

	// Bind the request body to the UserLoginForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Validate the login form
	errs := l.loginUseCase.ValidateUserForm(form)

	// Check if there are any errors
	if errs != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Process the login form
	user, processErr := l.loginUseCase.ProcessUser(form)

	// Check if there is an error processing the form
	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: processErr.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: processErr.Message,
			Data:    nil,
		})
	}

	// Generate a token
	token, err := l.authUseCase.GenerateToken(user.ID, enums.User)

	// Check if there is an error generating the token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error generating token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User Login Success",
		Data: dto.UserLoginResponseData{
			User:  dto.CurrentUser{}.FromModel(user),
			Token: token,
		},
	})
}

// VendorLogin is a function that handles the vendor login request.
// Endpoint: POST /api/v1/auth/vendor/login
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func (l *LoginController) VendorLogin(c echo.Context) error {
	// Create a new VendorLoginForm object
	form := new(dto.VendorLoginForm)

	// Bind the request body to the VendorLoginForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Validate the login form
	errs := l.loginUseCase.ValidateVendorForm(form)

	// Check if there are any errors
	if errs != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Process the login form
	vendor, processErr := l.loginUseCase.ProcessVendor(form)

	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: processErr.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: processErr.Message,
			Data:    nil,
		})
	}

	// Generate a token
	token, err := l.authUseCase.GenerateToken(vendor.ID, enums.Vendor)

	// Check if there is an error generating the token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Error generating token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Vendor Login Success",
		Data: dto.VendorLoginResponseData{
			Vendor: dto.CurrentVendor{}.FromModel(vendor),
			Token:  token,
		},
	})
}

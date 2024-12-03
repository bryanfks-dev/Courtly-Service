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
	LoginUseCase *usecases.LoginUseCase
	AuthUseCase  *usecases.AuthUseCase
}

// NewLoginController is a factory function that returns a new instance of the LoginController.
//
// l: The login use case.
// a: The auth use case.
//
// Returns a new instance of the LoginController.
func NewLoginController(l *usecases.LoginUseCase, a *usecases.AuthUseCase) *LoginController {
	return &LoginController{LoginUseCase: l, AuthUseCase: a}

}

// UserLogin is a function that handles the user login request.
// Endpoint: POST /auth/user/login
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
	errs := l.LoginUseCase.ValidateUserForm(form)

	// Check if there are any errors
	if errs != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Process the login form
	user, processErr := l.LoginUseCase.ProcessUser(form)

	// Check if there is an error processing the form
	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
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

	// Generate a token
	token, err := l.AuthUseCase.GenerateToken(user.ID, enums.User)

	// Check if there is an error generating the token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Error generating token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User Login Success",
		Data: dto.UserLoginResponseDTO{
			User:  dto.CurrentUserDTO{}.FromModel(user),
			Token: token,
		},
	})
}

// VendorLogin is a function that handles the vendor login request.
// Endpoint: POST /auth/vendor/login
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func (l *LoginController) VendorLogin(c echo.Context) error {
	// Create a new VendorLoginForm dto object
	form := new(dto.VendorLoginFormDTO)

	// Bind the request body to the VendorLoginForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Validate the login form
	errs := l.LoginUseCase.ValidateVendorForm(form)

	// Check if there are any errors
	if errs != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Process the login form
	vendor, processErr := l.LoginUseCase.ProcessVendor(form)

	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
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

	// Generate a token
	token, err := l.AuthUseCase.GenerateToken(vendor.ID, enums.Vendor)

	// Check if there is an error generating the token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Error generating token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Vendor Login Success",
		Data: dto.VendorLoginResponseDTO{
			Vendor: dto.CurrentVendorDTO{}.FromModel(vendor),
			Token:  token,
		},
	})
}

package controllers

import (
	"log"
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

// Login is a function that handles the login request.
// Endpoint: POST /api/v1/login
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func (l *LoginController) Login(c echo.Context) error {
	// Create a new LoginForm object
	form := new(dto.LoginForm)

	// Bind the request body to the LoginForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Validate the login form
	errs := l.loginUseCase.ValidateForm(form)

	// Check if there are any errors
	if errs != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Process the login form
	user, processErr := l.loginUseCase.Process(form)

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
	token, err := l.authUseCase.GenerateToken(user.ID)

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
		Message: "Login Success",
		Data: dto.LoginResponseData{
			User:  dto.CurrentUser{}.FromModel(user),
			Token: token,
		},
	})
}

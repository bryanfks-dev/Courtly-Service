package controllers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserController is a struct that defines the user controller.
type UserController struct {
	UserUseCase *usecases.UserUseCase
	AuthUseCase *usecases.AuthUseCase
}

// NewUserController is a factory function that returns a new instance of the UserController.
//
// u: The user use case.
// a: The auth use case.
//
// Returns a new instance of the UserController.
func NewUserController(u *usecases.UserUseCase, a *usecases.AuthUseCase) *UserController {
	return &UserController{
		UserUseCase: u,
		AuthUseCase: a,
	}
}

// GetCurrentUser is a handler function that returns the current user.
// Endpoint: GET /users/me
//
// c: The echo context.
//
// Returns an error if any.
func (u *UserController) GetCurrentUser(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current user
	user, err := u.UserUseCase.GetCurrentUser(cc.Token)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get current user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User retrieved successfully",
		Data: dto.CurrentUserResponseDTO{
			User: dto.CurrentUserDTO{}.FromModel(user),
		},
	})
}

// UpdateCurrentUserPassword is a handler function that updates the current user's password.
// Endpoint: PATCH /users/me/password
//
// c: The echo context.
//
// Returns an error if any.
func (u UserController) UpdateCurrentUserPassword(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Bind the form dto
	form := new(dto.ChangePasswordFormDTO)

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
	if err := u.UserUseCase.ValidateChangePasswordForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Update the password
	err := u.UserUseCase.ProcessChangePassword(cc.Token, form)

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
		Data:    nil,
	})
}

// UpdateCurrentUserUsername is a handler function that updates the current user's username.
// Endpoint: PATCH /users/me/username
//
// c: The echo context.
//
// Returns an error if any.
func (u *UserController) UpdateCurrentUserUsername(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Bind the form dto
	form := new(dto.ChangeUsernameFormDTO)

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
	if err := u.UserUseCase.ValidateChangeUsernameForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Update the username
	err := u.UserUseCase.ProcessChangeUsername(cc.Token, form)

	// Return an error if any
	if err != nil {
		if err.ClientError {
			return c.JSON(http.StatusForbidden, dto.ResponseDTO{
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
		Message: "Username updated successfully",
		Data:    nil,
	})
}

package controllers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"
	"strconv"

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

	user, err := u.UserUseCase.GetCurrentUser(cc.Token)

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

// GetPublicUser is a handler function that returns the user with the given ID, with restrict information.
// Endpoint: GET /users/:id
//
// c: The echo context.
//
// Returns an error if any.
func (u *UserController) GetPublicUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	id := c.Param("id")

	// Convert the ID to an integer
	userID, err := strconv.Atoi(id)

	// Return an error if the ID is not a valid integer
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid user ID",
			Data:    nil,
		})
	}

	// Get the user with the given ID
	user, processErr := u.UserUseCase.GetUserUsingID(uint(userID))

	// Return an error if the user does not exist
	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
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

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User retrieved successfully",
		Data: dto.PublicUserResponseDTO{
			User: dto.PublicUserDTO{}.FromModel(user),
		},
	})
}

// UpdateUserPassword is a handler function that updates the current user's password.
// Endpoint: PATCH /users/me/password
//
// c: The echo context.
//
// Returns an error if any.
func (u UserController) UpdateUserPassword(c echo.Context) error {
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
	if err := u.UserUseCase.ValidateChangePasswordForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Update the password
	user, err := u.UserUseCase.ProcessChangePassword(cc.Token, form)

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
		Data: dto.CurrentUserResponseDTO{
			User: dto.CurrentUserDTO{}.FromModel(user),
		},
	})
}

// UpdateUserUsername is a handler function that updates the current user's username.
// Endpoint: PATCH /users/me/username
//
// c: The echo context.
//
// Returns an error if any.
func (u *UserController) UpdateUserUsername(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Bind the form data
	form := new(dto.ChangeUsernameForm)

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
	user, err := u.UserUseCase.ProcessChangeUsername(cc.Token, form)

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
		Data: dto.CurrentUserResponseDTO{
			User: dto.CurrentUserDTO{}.FromModel(user),
		},
	})
}

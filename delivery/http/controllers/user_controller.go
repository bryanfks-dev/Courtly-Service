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
	userUseCase           *usecases.UserUseCase
	changePasswordUseCase *usecases.ChangeUserPasswordUseCase
	changeUsernameUseCase *usecases.ChangeUserUsernameUseCase
	authUseCase           *usecases.AuthUseCase
}

// NewUserController is a factory function that returns a new instance of the UserController.
//
// u: The user use case.
// cp: The change password use case.
// cu: The change username use case.
//
// Returns a new instance of the UserController.
func NewUserController(u *usecases.UserUseCase, cp *usecases.ChangeUserPasswordUseCase, cu *usecases.ChangeUserUsernameUseCase, a *usecases.AuthUseCase) *UserController {
	return &UserController{
		userUseCase:           u,
		changePasswordUseCase: cp,
		changeUsernameUseCase: cu,
		authUseCase:           a,
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

	// Decode the token
	claims := u.authUseCase.DecodeToken(cc.Token)

	// Get the user from the database
	user, err := u.userUseCase.GetUserByID(claims.Id)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "An error occurred while getting the user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User retrieved successfully",
		Data: dto.CurrentUserResponseData{
			User: dto.CurrentUser{}.FromModel(user),
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
		return c.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid user ID",
			Data:    nil,
		})
	}

	// Get the user with the given ID
	user, err := u.userUseCase.GetUserByID(uint(userID))

	// Return an error if the user does not exist
	if err != nil {
		return c.JSON(404, dto.Response{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
	}

	return c.JSON(200, dto.Response{
		Success: true,
		Message: "User retrieved successfully",
		Data: dto.PublicUserResponseData{
			User: dto.PublicUser{}.FromModel(*user),
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

	// Decode the token
	claims := u.authUseCase.DecodeToken(cc.Token)

	// Bind the form data
	form := new(dto.ChangePasswordForm)

	// Return an error if the form data is invalid
	if err := c.Bind(form); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Validate the form data
	if err := u.changePasswordUseCase.ValidateForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Update the password
	user, err := u.changePasswordUseCase.Process(claims.Id, form)

	// Return an error if any
	if err != nil {
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
		Message: "Password updated successfully",
		Data: dto.CurrentUserResponseData{
			User: dto.CurrentUser{}.FromModel(user),
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

	// Decode the token
	claims := u.authUseCase.DecodeToken(cc.Token)

	// Bind the form data
	form := new(dto.ChangeUsernameForm)

	// Return an error if the form data is invalid
	if err := c.Bind(form); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Validate the form data
	if err := u.changeUsernameUseCase.ValidateForm(form); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err,
			Data:    nil,
		})
	}

	// Update the username
	user, err := u.changeUsernameUseCase.Process(claims.Id, form)

	// Return an error if any
	if err != nil {
		if err.ClientError {
			return c.JSON(http.StatusForbidden, dto.Response{
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
		Message: "Username updated successfully",
		Data: dto.CurrentUserResponseData{
			User: dto.CurrentUser{}.FromModel(user),
		},
	})
}

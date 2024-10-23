package controllers

import (
	"main/domain/usecases"
	"main/internal/dto"
	"main/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetCurrentUser is a handler function that returns the current user.
// Endpoint: GET /users/me
//
// c: The echo context.
//
// Returns an error if any.
func GetCurrentUser(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Decode the token
	claims := usecases.DecodeToken(cc.Token)

	// Get the user from the database
	user, err := repository.GetUserByID(claims["id"].(uint))

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "An error occurred while getting the user",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		StatusCode: http.StatusOK,
		Message:    "User retrieved successfully",
		Data: dto.CurrentUserResponseData{
			User: user,
		},
	})
}

// GetPublicUser is a handler function that returns the user with the given ID, with restrict information.
// Endpoint: GET /users/:id
//
// c: The echo context.
//
// Returns an error if any.
func GetPublicUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	id := c.Param("id")

	// Convert the ID to an integer
	userID, err := strconv.Atoi(id)

	// Return an error if the ID is not a valid integer
	if err != nil {
		return c.JSON(400, dto.Response{
			StatusCode: 400,
			Message:    "Invalid user ID",
			Data:       nil,
		})
	}

	// Get the user with the given ID
	user, err := repository.GetUserByID(uint(userID))

	// Return an error if the user does not exist
	if err != nil {
		return c.JSON(404, dto.Response{
			StatusCode: 404,
			Message:    "User not found",
			Data:       nil,
		})
	}

	return c.JSON(200, dto.Response{
		StatusCode: 200,
		Message:    "Success get user",
		Data:       user,
	})
}

package handlers

import (
	"main/domain/usecases"
	"main/internal/dto"
	"main/internal/repository"
	"net/http"

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
		Data:       user,
	})
}

package handlers

import (
	"main/internal/dto"
	"main/internal/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetUser is a handler function that returns the user with the given ID.
// Endpoint: GET /users/:id
//
// c: The echo context.
//
// Returns an error if any.
func GetUser(c echo.Context) error {
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

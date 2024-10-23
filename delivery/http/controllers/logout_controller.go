package controllers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Logout is a handler that logs out a user
// by blacklisting the token used to authenticate the user.
// Endpoint: POST /logout
//
// c: echo.Context
//
// Returns an error response if there is an error, otherwise a success response.
func Logout(c echo.Context) error {
	cc := c.(*dto.CustomContext)

	// Blacklist token
	err := usecases.BlacklistToken(cc.Token)

	// Check if there was an error blacklisting the token
	if err != nil {
		log.Fatal("Error blacklisting token: ", err)

		return c.JSON(http.StatusInternalServerError, dto.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Could not blacklist token",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		StatusCode: http.StatusOK,
		Message:    "Successfully logged out",
		Data:       nil,
	})
}

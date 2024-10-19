package handlers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"main/pkg/utils"
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
	// Extract token from request
	tokenString, err := usecases.ExtractToken(c)

	// Check if there was an error extracting the token
	if err != nil {
		// In this case, there is no need to log the extraction
		// error, because the extraction error caused by the
		// user request

		return c.JSON(http.StatusUnauthorized, dto.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    utils.ToUpperFirst(err.Error()),
			Data:       nil,
		})
	}

	// Validate token
	token, valid := usecases.VerifyToken(tokenString)

	// Check if token is valid
	if !valid {
		// In this case, there is no need to log the verify token
		// error, because the verifying error caused by the
		// user request

		return c.JSON(http.StatusUnauthorized, dto.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid token",
			Data:       nil,
		})
	}

	// Blacklist token
	err = usecases.BlacklistToken(token)

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

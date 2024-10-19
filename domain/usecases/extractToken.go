package usecases

import (
	"errors"
	"main/pkg/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

// ExtractToken is a usecase that extracts the token from the request header
//
// c: the echo.Context object
//
// Returns the token string and an error if there is any
func ExtractToken(c echo.Context) (string, error) {
	// Extract the token from the request header
	tokenString := c.Request().Header.Get("Authorization")

	// Check if the token is blank
	if utils.IsBlank(tokenString) {
		return "", errors.New("token is missing")
	}

	// Sanitize the tokenString
	tokenString = strings.TrimSpace(tokenString)

	token := strings.Split(tokenString, " ")

	// Check if the token is in the correct format
	if token[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	return token[1], nil
}

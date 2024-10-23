package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// CustomContext is a custom context that extends the echo.Context
// with additional fields.
type CustomContext struct {
	echo.Context

	// Token is the JWT token extracted from the request.
	Token *jwt.Token
}

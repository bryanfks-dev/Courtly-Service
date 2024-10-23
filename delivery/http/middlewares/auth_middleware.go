package middlewares

import (
	"main/domain/usecases"
	"main/internal/dto"
	"main/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
// and if not, it stops the request and returns a 401 status code.
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the token from the request
		tokenString, err := usecases.ExtractToken(c)

		// If there was an error extracting the token, return the error
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    utils.ToUpperFirst(err.Error()),
				Data:       nil,
			})
		}

		// Validate the token
		token, valid := usecases.VerifyToken(tokenString)

		// If the token is not valid, return a 401 status code
		if !valid {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token",
				Data:       nil,
			})
		}

		// Create a custom context with the token
		cc := &dto.CustomContext{
			Context: c,
			Token:   token,
		}

		// Call the next handler
		return next(cc)
	}
}

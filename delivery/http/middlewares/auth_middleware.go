package middlewares

import (
	"main/domain/usecases"
	"main/internal/dto"
	"main/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
type AuthMiddleware struct {
	authUseCase *usecases.AuthUseCase
}

// NewAuthMiddleware is a factory function that returns a new instance of the AuthMiddleware
// Returns a new instance of the AuthMiddleware
func NewAuthMiddleware(a *usecases.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{authUseCase: a}
}

// Shield is a middleware that checks if the user is authenticated
//
// next: The next handler function
//
// Returns an error if any
func (a *AuthMiddleware) Shield(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the token from the request
		tokenString, err := a.authUseCase.ExtractToken(c)

		// If there was an error extracting the token, return the error
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: utils.ToUpperFirst(err.Error()),
				Data:    nil,
			})
		}

		// Validate the token
		token, valid := a.authUseCase.VerifyToken(tokenString)

		// If the token is not valid, return a 401 status code
		if !valid {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "Invalid token",
				Data:    nil,
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

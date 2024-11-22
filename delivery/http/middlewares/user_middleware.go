package middlewares

import (
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserMiddleware is a middleware that checks if the token is users, not other client type
type UserMiddleware struct {
	authUseCase *usecases.AuthUseCase
}

// NewUserMiddleware is a factory function that returns a new instance of the UserMiddleware
//
// a: The auth use case
//
// Returns a new instance of the UserMiddleware
func NewUserMiddleware(a *usecases.AuthUseCase) *UserMiddleware {
	return &UserMiddleware{authUseCase: a}
}

// Shield is a middleware that checks if the user is authenticated
//
// next: The next handler function
//
// Returns an error if any
func (u *UserMiddleware) Shield(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get custom context
		cc := c.(*dto.CustomContext)

		// Check if token is users, not other client type
		if u.authUseCase.DecodeToken(cc.Token).ClientType != enums.User {
			return c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Success: false,
				Message: "Invalid client type for this endpoint",
				Data:    nil,
			})
		}

		// Call the next handler
		return next(cc)
	}
}

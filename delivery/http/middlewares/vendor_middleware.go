package middlewares

import (
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VendorMiddleware is a middleware that checks if the token is users, not other client type
type VendorMiddleware struct {
	authUseCase *usecases.AuthUseCase
}

// NewVendorMiddleware is a factory function that returns a new instance of the VendorMiddleware
//
// a: The auth use case
//
// Returns a new instance of the VendorMiddleware
func NewVendorMiddleware(a *usecases.AuthUseCase) *VendorMiddleware {
	return &VendorMiddleware{authUseCase: a}
}

// Shield is a middleware that checks if the user is authenticated
//
// next: The next handler function
//
// Returns an error if any
func (v *VendorMiddleware) Shield(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get custom context
		cc := c.(*dto.CustomContext)

		// Check if token is users, not other client type
		if v.authUseCase.DecodeToken(cc.Token).ClientType != enums.Vendor {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "Invalid client type for this endpoint",
				Data:    nil,
			})
		}

		// Call the next handler
		return next(cc)
	}
}

package middlewares

import (
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BlacklistTokenMiddleware is a middleware that checks if the token is blacklisted.
type BlacklistTokenMiddleware struct {
	blacklistTokenUseCase *usecases.BlacklistedTokenUseCase
}

// NewBlacklistTokenMiddleware is a factory function that returns a new instance of the BlacklistTokenMiddleware.
//
// b: The blacklisted token use case.
//
// Returns a new instance of the BlacklistTokenMiddleware.
func NewBlacklistTokenMiddleware(b *usecases.BlacklistedTokenUseCase) *BlacklistTokenMiddleware {
	return &BlacklistTokenMiddleware{blacklistTokenUseCase: b}
}

// Shield is a function that checks if the token is blacklisted.
//
// next: The next handler function.
//
// Returns an error if any.
func (b *BlacklistTokenMiddleware) Shield(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the token from the request
		cc := c.(*dto.CustomContext)

		// Check if the token is blacklisted
		if b.blacklistTokenUseCase.IsBlacklistedToken(cc.Token.Raw) {
			return c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "Token is blacklisted",
				Data:    nil,
			})
		}

		return next(cc)
	}
}

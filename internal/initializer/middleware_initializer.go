package initializer

import "main/delivery/http/middlewares"

// Middlewares is a struct that holds all the middlewares.
type Middlewares struct {
	AuthMiddleware             *middlewares.AuthMiddleware
	BlacklistedTokenMiddleware *middlewares.BlacklistTokenMiddleware
	UserMiddleware             *middlewares.UserMiddleware
	VendorMiddleware           *middlewares.VendorMiddleware
}

// InitMiddlewares is a function that initializes all the middlewares.
//
// usecase: Instance of UseCases
//
// Returns an instance of Middlewares.
func InitMiddlewares(usecase *UseCases) *Middlewares {
	return &Middlewares{
		AuthMiddleware:             middlewares.NewAuthMiddleware(usecase.AuthUseCase),
		BlacklistedTokenMiddleware: middlewares.NewBlacklistTokenMiddleware(usecase.BlacklistedTokenUseCase),
		UserMiddleware:             middlewares.NewUserMiddleware(usecase.AuthUseCase),
		VendorMiddleware:           middlewares.NewVendorMiddleware(usecase.AuthUseCase),
	}
}

package server

import (
	"main/delivery/http/middlewares"
)

var (
	// authMiddleware is a middleware that checks if a user is authenticated.
	authMiddleware *middlewares.AuthMiddleware

	// blacklistedTokenMiddleware is a middleware that checks if a token is blacklisted.
	blacklistedTokenMiddleware *middlewares.BlacklistTokenMiddleware
)

// initMiddleware is a function that initializes the middlewares.
func initMiddleware() {
	authMiddleware = middlewares.NewAuthMiddleware(authUseCase)
	blacklistedTokenMiddleware = middlewares.NewBlacklistTokenMiddleware(blacklistedTokenUseCase)
}

package server

import (
	"main/core/config"
	"main/core/constants"
	"main/delivery/http/router"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewServer is a factory function that returns a new instance of the echo.Echo server
// with the given configuration.
//
// Returns the echo.Echo server instance and an error if any.
func NewServer() (*echo.Echo, error) {
	e := echo.New()

	// Initialize repositories
	initRepositories()

	// Initialize use cases
	initUseCases()

	// Initialize controllers
	initControllers()

	// Initialize middlewares
	initMiddleware()

	e.Use(middleware.CORS())

	// Register static files
	e.Static(router.UserProfiles, constants.PATH_TO_USER_PROFILE_PICTURES)

	// Register prefix endpoint
	prefix := e.Group("/api/v1")

	// Endpoint list
	prefix.POST("/register", registerController.Register)

	prefix.POST("/login", loginController.Login)

	prefix.POST("/logout", logoutController.Logout, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	prefix.GET("/users/me", userController.GetCurrentUser, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	prefix.GET("/users/:id", userController.GetPublicUser, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	prefix.PATCH("/users/me/password", userController.UpdateUserPassword, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	prefix.PATCH("/users/me/username", userController.UpdateUserUsername, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	return e, e.Start(":" + strconv.Itoa(config.ServerConfig.Port))
}

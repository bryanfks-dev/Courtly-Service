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
	// Auth endpoints
	authPrefix := prefix.Group("/auth")

	authPrefix.POST("/register", registerController.Register)
	authPrefix.POST("/login", loginController.Login)
	authPrefix.POST("/logout", logoutController.Logout, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)
	authPrefix.POST("/verify-password", verifyPasswordController.VerifyPassword, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	// User endpoints
	userPrefix := prefix.Group("/users")

	// Current user endpoints
	currentUserPrefix := userPrefix.Group("/me")

	currentUserPrefix.GET("/", userController.GetCurrentUser, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)
	currentUserPrefix.PATCH("/username", userController.UpdateUserUsername, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)
	currentUserPrefix.PATCH("/password", userController.UpdateUserPassword, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	prefix.GET("/:id", userController.GetPublicUser, authMiddleware.Shield, blacklistedTokenMiddleware.Shield)

	return e, e.Start(":" + strconv.Itoa(config.ServerConfig.Port))
}

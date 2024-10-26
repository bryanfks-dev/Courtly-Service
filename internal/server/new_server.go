package server

import (
	"main/core/config"
	"main/core/constants"
	"main/delivery/http/router"
	"main/internal/initializer"
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

	// Repository initialization
	r := initializer.InitRepositories()

	// Usecase initialization
	u := initializer.InitUseCases(r)

	// Controller initialization
	c := initializer.InitControllers(u)

	// Middleware initialization
	m := initializer.InitMiddlewares(u)

	// Register middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Register static files
	e.Static(router.UserProfiles, constants.PATH_TO_USER_PROFILE_PICTURES)

	// Register prefix endpoint
	prefix := e.Group("/api/v1")

	// Register routes
	// Auth endpoints
	authPrefix := prefix.Group("/auth")

	authPrefix.POST("/register", c.RegisterController.Register)
	authPrefix.POST("/login", c.LoginController.Login)
	authPrefix.POST("/logout", c.LogoutController.Logout, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)
	authPrefix.POST("/verify-password", c.VerifyPasswordController.VerifyPassword, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)

	// User endpoints
	userPrefix := prefix.Group("/users")

	prefix.GET("/:id", c.UserController.GetPublicUser, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)

	// Current user endpoints
	currentUserPrefix := userPrefix.Group("/me")

	currentUserPrefix.GET("", c.UserController.GetCurrentUser, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)
	currentUserPrefix.PATCH("/username", c.UserController.UpdateUserUsername, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)
	currentUserPrefix.PATCH("/password", c.UserController.UpdateUserPassword, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)

	// User endpoints

	return e, e.Start(":" + strconv.Itoa(config.ServerConfig.Port))
}

package server

import (
	"main/core/config"
	"main/core/constants"
	"main/delivery/http/controllers"
	"main/delivery/http/middlewares"
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

	e.Use(middleware.CORS())

	// Register static files
	e.Static(router.UserProfiles, constants.PATH_TO_USER_PROFILE_PICTURES)

	// Register prefix endpoint
	prefix := e.Group("/api/v1")

	// Endpoint list
	prefix.POST("/register", controllers.Register)
	prefix.POST("/login", controllers.Login)
	prefix.POST("/logout", controllers.Logout, middlewares.AuthMiddleware)
	prefix.GET("/users/me", controllers.GetCurrentUser, middlewares.AuthMiddleware)
	prefix.GET("/users/:id", controllers.GetPublicUser, middlewares.AuthMiddleware)

	return e, e.Start(":" + strconv.Itoa(config.ServerConfig.Port))
}

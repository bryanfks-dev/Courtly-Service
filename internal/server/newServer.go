package server

import (
	"main/core/config"
	"main/delivery/http/controllers"
	"main/delivery/http/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewServer is a factory function that returns a new instance of the echo.Echo server
// with the given configuration.
//
// serverConfig: The configuration of the server.
//
// Returns the echo.Echo server instance and an error if any.
func NewServer(serverConfig config.Server) (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.CORS())

	// Register prefix endpoint
	prefix := e.Group("/api/v1")

	// Endpoint list
	prefix.POST("/register", controllers.Register)
	prefix.POST("/login", controllers.Login)
	prefix.POST("/logout", controllers.Logout, middlewares.AuthMiddleware)
	prefix.GET("/users/me", controllers.GetCurrentUser, middlewares.AuthMiddleware)

	return e, e.Start(":" + strconv.Itoa(serverConfig.Port))
}

package server

import (
	"main/core/config"
	"main/internal/handlers"
	"main/internal/middlewares"
	"strconv"

	"github.com/labstack/echo-jwt/v4"
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

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTConfig.Secret),
	}))

	// Register prefix endpoint
	prefix := e.Group("/api/v1")

	// Endpoint list
	prefix.POST("/register", handlers.Register)
	prefix.POST("/login", handlers.Login)
	prefix.POST("/logout", handlers.Logout, middlewares.AuthMiddleware)
	prefix.GET("/users/me", handlers.GetCurrentUser, middlewares.AuthMiddleware)

	return e, e.Start(":" + strconv.Itoa(serverConfig.Port))
}

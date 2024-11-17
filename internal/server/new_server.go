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

	// User Auth endpoints
	userAuthPrefix := authPrefix.Group("/user")

	userAuthPrefix.POST("/register", c.RegisterController.UserRegister)
	userAuthPrefix.POST("/login", c.LoginController.UserLogin)
	userAuthPrefix.POST("/verify-password", c.VerifyPasswordController.UserVerifyPassword, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)
	userAuthPrefix.POST("/logout", c.LogoutController.UserLogout, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.VendorMiddleware.Shield)

	// Vendor Auth endpoints
	vendorAuthPrefix := authPrefix.Group("/vendor")

	vendorAuthPrefix.POST("/login", c.LoginController.VendorLogin)
	vendorAuthPrefix.POST("/logout", c.LogoutController.VendorLogout, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.VendorMiddleware.Shield)

	// Users endpoints
	userPrefix := prefix.Group("/users")

	userPrefix.GET("/:id", c.UserController.GetPublicUser, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)

	// Current user endpoints
	currentUserPrefix := userPrefix.Group("/me", m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.UserMiddleware.Shield)

	currentUserPrefix.GET("", c.UserController.GetCurrentUser)
	currentUserPrefix.PATCH("/username", c.UserController.UpdateUserUsername)
	currentUserPrefix.PATCH("/password", c.UserController.UpdateUserPassword)

	// Vendors endpoints
	vendorPrefix := prefix.Group("/vendors")

	// Current vendor endpoints
	currentVendorPrefix := vendorPrefix.Group("/me", m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.VendorMiddleware.Shield)

	currentVendorPrefix.GET("", c.VendorController.GetCurrentVendor)
	currentVendorPrefix.PATCH("/password", c.VendorController.UpdateVendorPassword)

	// Courts endpoints
	currentVendorCourtsPrefix := currentVendorPrefix.Group("/courts")

	currentVendorCourtsPrefix.GET("/:type", c.CourtController.GetCurrentVendorCourtType)

	return e, e.Start(":" + strconv.Itoa(config.ServerConfig.Port))
}

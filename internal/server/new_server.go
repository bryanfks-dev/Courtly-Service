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

	e.Static(router.CourtImages, constants.PATH_TO_COURT_IMAGES)

	// Register prefix endpoint
	prefix := e.Group("/api/v1")

	// Fees endpoint
	prefix.GET("/fees", c.FeesController.GetFees, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)

	// Register routes
	// Auth endpoints
	authPrefix := prefix.Group("/auth")

	// User Auth endpoints
	userAuthPrefix := authPrefix.Group("/user")

	userAuthPrefix.POST("/register", c.RegisterController.UserRegister)
	userAuthPrefix.POST("/login", c.LoginController.UserLogin)
	userAuthPrefix.POST("/verify-password", c.VerifyPasswordController.UserVerifyPassword, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)
	userAuthPrefix.POST("/logout", c.LogoutController.UserLogout, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.UserMiddleware.Shield)

	// Vendor Auth endpoints
	vendorAuthPrefix := authPrefix.Group("/vendor")

	vendorAuthPrefix.POST("/login", c.LoginController.VendorLogin)
	vendorAuthPrefix.POST("/verify-password", c.VerifyPasswordController.VendorVerifyPassword, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield)
	vendorAuthPrefix.POST("/logout", c.LogoutController.VendorLogout, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.VendorMiddleware.Shield)

	// Users endpoints
	userPrefix := prefix.Group("/users")

	// Current user endpoints
	currentUserPrefix := userPrefix.Group("/me", m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.UserMiddleware.Shield)

	currentUserPrefix.GET("", c.UserController.GetCurrentUser)

	currentUserPrefix.PATCH("/username", c.UserController.UpdateCurrentUserUsername)

	currentUserPrefix.PATCH("/password", c.UserController.UpdateCurrentUserPassword)

	currentUserPrefix.PATCH("/profile-picture", c.UserController.UpdateCurrentUserProfilePicture)

	// Vendors endpoints
	vendorPrefix := prefix.Group("/vendors")

	// Current vendor endpoints
	currentVendorPrefix := vendorPrefix.Group("/me", m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.VendorMiddleware.Shield)

	currentVendorPrefix.GET("", c.VendorController.GetCurrentVendor)
	currentVendorPrefix.PATCH("/password", c.VendorController.UpdateCurrentVendorPassword)

	// Current user orders endpoints
	currentUserOrdersPrefix := currentUserPrefix.Group("/orders")

	currentUserOrdersPrefix.GET("", c.OrderController.GetCurrentUserOrders)

	currentUserOrdersPrefix.POST("", c.OrderController.CreateOrder)

	currentUserOrdersPrefix.GET("/:id", c.OrderController.GetCurrentUserOrderDetail)

	// Current vendor orders endpoints
	currentVendorOrdersPrefix := currentVendorPrefix.Group("/orders")

	currentVendorOrdersPrefix.GET("", c.OrderController.GetCurrentVendorOrders)

	currentVendorOrdersPrefix.GET("/stats", c.OrderController.GetCurrentVendorOrdersStats)

	currentVendorOrdersPrefix.GET("/:id", c.OrderController.GetCurrentVendorOrderDetail)

	// Courts endpoints
	courtPrefix := prefix.Group("/courts")

	courtPrefix.GET("", c.CourtController.GetCourts)

	// Vendor courts endpoints
	vendorCourtsPrefix := vendorPrefix.Group("/:id/courts", m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.UserMiddleware.Shield)

	vendorTypeCourtsPrefix := vendorCourtsPrefix.Group("/:type")

	vendorTypeCourtsPrefix.GET("", c.CourtController.GetVendorCourtsUsingCourtType)

	vendorTypeCourtsPrefix.GET("/bookings", c.CourtController.GetCourtBookings)

	// Current vendor courts endpoints
	currentVendorCourtsPrefix := currentVendorPrefix.Group("/courts")

	currentVendorCourtsPrefix.GET("/stats", c.CourtController.GetCurrentVendorCourtStats)

	currentVendorCourtsPrefix.DELETE("", c.CourtController.DeleteCourts)

	// Current vendor courts types endpoints
	currentVendorCourtsTypePrefix := currentVendorCourtsPrefix.Group("/:type")

	currentVendorCourtsTypePrefix.GET("/bookings", c.CourtController.GetCurrentVendorCourtBookings)

	currentVendorCourtsTypePrefix.GET("", c.CourtController.GetCurrentVendorCourtsUsingCourtType)

	currentVendorCourtsTypePrefix.POST("", c.CourtController.AddCourt)

	currentVendorCourtsTypePrefix.PUT("", c.CourtController.UpdateCourtUsingCourtType)

	currentVendorCourtsTypePrefix.POST("/new", c.CourtController.CreateNewCourt)

	// Reviews endpoints
	vendorTypeCourtsPrefix.GET("/reviews", c.ReviewController.GetCourtTypeReviews)

	vendorTypeCourtsPrefix.POST("/reviews", c.ReviewController.CreateReview, m.AuthMiddleware.Shield, m.BlacklistedTokenMiddleware.Shield, m.UserMiddleware.Shield)

	currentVendorPrefix.GET("/reviews", c.ReviewController.GetCurrentVendorReviews)

	// Midtrans endpoints
	midtransPrefix := e.Group("/midtrans")

	midtransPrefix.POST("/payment-callback", c.MidtransController.PaymentCallback)

	return e, e.Start(":" + strconv.Itoa(config.ServerConfig.Port))
}

package initializer

import "main/delivery/http/controllers"

// Controllers is a struct that holds all the controllers.
type Controllers struct {
	LoginController          *controllers.LoginController
	RegisterController       *controllers.RegisterController
	LogoutController         *controllers.LogoutController
	VerifyPasswordController *controllers.VerifyPasswordController
	UserController           *controllers.UserController
	VendorController         *controllers.VendorController
	CourtController          *controllers.CourtController
	ReviewController         *controllers.ReviewController
	BookingController        *controllers.BookingController
	OrderController          *controllers.OrderController
}

// InitControllers is a function that initializes all the controllers.
//
// usecase: Instance of UseCases
//
// Returns an instance of Controllers.
func InitControllers(usecase *UseCases) *Controllers {
	return &Controllers{
		LoginController:          controllers.NewLoginController(usecase.LoginUseCase, usecase.AuthUseCase),
		RegisterController:       controllers.NewRegisterController(usecase.RegisterUseCase),
		LogoutController:         controllers.NewLogoutController(usecase.LogoutUseCase),
		VerifyPasswordController: controllers.NewVerifyPasswordController(usecase.VerifyPasswordUseCase),
		UserController:           controllers.NewUserController(usecase.UserUseCase, usecase.AuthUseCase),
		VendorController:         controllers.NewVendorController(usecase.VendorUseCase),
		CourtController:          controllers.NewCourtController(usecase.CourtUseCase),
		ReviewController:         controllers.NewReviewController(usecase.ReviewUseCase),
		BookingController:        controllers.NewBookingController(usecase.BookingUseCase, usecase.OrderUseCase),
		OrderController:          controllers.NewOrderController(usecase.OrderUseCase, usecase.ReviewUseCase),
	}
}

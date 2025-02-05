package initializer

import "main/delivery/http/controllers"

// Controllers is a struct that holds all the controllers.
type Controllers struct {
	FeesController           *controllers.FeesController
	LoginController          *controllers.LoginController
	RegisterController       *controllers.RegisterController
	LogoutController         *controllers.LogoutController
	VerifyPasswordController *controllers.VerifyPasswordController
	UserController           *controllers.UserController
	VendorController         *controllers.VendorController
	CourtController          *controllers.CourtController
	ReviewController         *controllers.ReviewController
	OrderController          *controllers.OrderController
	AdvertisementController  *controllers.AdvertisementController
	MidtransController       *controllers.MidtransController
}

// InitControllers is a function that initializes all the controllers.
//
// usecase: Instance of UseCases
//
// Returns an instance of Controllers.
func InitControllers(usecase *UseCases) *Controllers {
	return &Controllers{
		FeesController:           controllers.NewFeesController(),
		LoginController:          controllers.NewLoginController(usecase.LoginUseCase, usecase.AuthUseCase),
		RegisterController:       controllers.NewRegisterController(usecase.RegisterUseCase),
		LogoutController:         controllers.NewLogoutController(usecase.LogoutUseCase),
		VerifyPasswordController: controllers.NewVerifyPasswordController(usecase.VerifyPasswordUseCase),
		UserController:           controllers.NewUserController(usecase.UserUseCase, usecase.AuthUseCase),
		VendorController:         controllers.NewVendorController(usecase.VendorUseCase),
		CourtController:          controllers.NewCourtController(usecase.CourtUseCase, usecase.BookingUseCase),
		ReviewController:         controllers.NewReviewController(usecase.ReviewUseCase),
		OrderController:          controllers.NewOrderController(usecase.OrderUseCase, usecase.ReviewUseCase),
		AdvertisementController:  controllers.NewAdvertisementController(usecase.AdvertisementUseCase),
		MidtransController:       controllers.NewMidtransController(),
	}
}

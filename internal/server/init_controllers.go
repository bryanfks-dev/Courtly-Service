package server

import (
	"main/delivery/http/controllers"
)

var (
	// registerController is a variable that holds the register controller.
	registerController controllers.RegisterController

	// loginController is a variable that holds the login controller.
	loginController controllers.LoginController

	// logoutController is a variable that holds the logout controller.
	logoutController controllers.LogoutController

	// userController is a variable that holds the user controller.
	userController controllers.UserController
)

// initControllers is a function that initializes the controllers.
//
// Returns void.
func initControllers() {
	registerController = controllers.NewRegisterController(
		registerUseCase,
	)

	loginController = controllers.NewLoginController(
		loginUseCase,
		authUseCase,
	)

	logoutController = controllers.NewLogoutController(
		logoutUseCase,
	)

	userController = controllers.NewUserController(
		userUseCases,
		changeUserPasswordUseCase,
		changeUserUsernameUseCase,
		authUseCase,
	)
}

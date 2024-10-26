package server

import "main/domain/usecases"

var (
	// authUseCase is a variable that holds the auth use case.
	authUseCase *usecases.AuthUseCase

	// registerUseCase is a variable that holds the register use case.
	registerUseCase *usecases.RegisterUseCase

	// loginUseCase is a variable that holds the login use case.
	loginUseCase *usecases.LoginUseCase

	// logoutUseCase is a variable that holds the logout use case.
	logoutUseCase *usecases.LogoutUseCase

	// userUseCases is a variable that holds the user use cases.
	userUseCases *usecases.UserUseCase

	// changeUserUsernameUseCase is a variable that holds the change username use case.
	changeUserUsernameUseCase *usecases.ChangeUserUsernameUseCase

	// changeUserPasswordUseCase is a variable that holds the change password use case.
	changeUserPasswordUseCase *usecases.ChangeUserPasswordUseCase

	// blacklistedTokenUseCase is a variable that holds the blacklisted token use case.
	blacklistedTokenUseCase *usecases.BlacklistedTokenUseCase
)

// initUseCases is a function that initializes the use cases.
//
// Returns void.
func initUseCases() {
	authUseCase = usecases.NewAuthUseCase()

	userUseCases = usecases.NewUserUseCase(userRepository)

	changeUserUsernameUseCase = usecases.NewChangeUserUsernameUseCase(userRepository)

	changeUserPasswordUseCase = usecases.NewChangeUserPasswordUseCase(userRepository)

	registerUseCase = usecases.NewRegisterUseCase(userRepository)

	loginUseCase = usecases.NewLoginUseCase(userRepository, authUseCase)

	logoutUseCase = usecases.NewLogoutUseCase(blacklistedTokenRepository, authUseCase)

	blacklistedTokenUseCase = usecases.NewBlacklistedTokenUseCase(blacklistedTokenRepository)
}

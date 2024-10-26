package initializer

import "main/domain/usecases"

// UseCases is a struct that holds all the use cases.
type UseCases struct {
	AuthUseCase               *usecases.AuthUseCase
	RegisterUseCase           *usecases.RegisterUseCase
	LoginUseCase              *usecases.LoginUseCase
	LogoutUseCase             *usecases.LogoutUseCase
	ChangeUserPasswordUseCase *usecases.ChangeUserPasswordUseCase
	ChangeUserUsernameUseCase *usecases.ChangeUserUsernameUseCase
	UserUseCase               *usecases.UserUseCase
	BlacklistedTokenUseCase   *usecases.BlacklistedTokenUseCase
}

// InitUseCases is a function that initializes all the use cases.
//
// repos: The repositories.
//
// Returns a pointer to the UseCases struct.
func InitUseCases(repos *Repositories) *UseCases {
	u := &UseCases{}

	u.AuthUseCase = usecases.NewAuthUseCase()

	u.RegisterUseCase = usecases.NewRegisterUseCase(repos.UserRepository)

	u.LoginUseCase = usecases.NewLoginUseCase(repos.UserRepository, u.AuthUseCase)

	u.LogoutUseCase = usecases.NewLogoutUseCase(repos.BlacklistedTokenRepository, u.AuthUseCase)

	u.ChangeUserPasswordUseCase = usecases.NewChangeUserPasswordUseCase(repos.UserRepository)

	u.ChangeUserUsernameUseCase = usecases.NewChangeUserUsernameUseCase(repos.UserRepository)

	u.UserUseCase = usecases.NewUserUseCase(repos.UserRepository)

	u.BlacklistedTokenUseCase = usecases.NewBlacklistedTokenUseCase(repos.BlacklistedTokenRepository)

	return u
}

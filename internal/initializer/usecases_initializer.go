package initializer

import "main/domain/usecases"

// UseCases is a struct that holds all the use cases.
type UseCases struct {
	AuthUseCase               *usecases.AuthUseCase
	VerifyPasswordUseCase     *usecases.VerifyPasswordUseCase
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

	u.VerifyPasswordUseCase = usecases.NewVerifyPasswordUseCase(u.AuthUseCase, repos.UserRepository)

	u.RegisterUseCase = usecases.NewRegisterUseCase(u.AuthUseCase, repos.UserRepository)

	u.LoginUseCase = usecases.NewLoginUseCase(u.AuthUseCase, repos.UserRepository, repos.VendorRepository)

	u.LogoutUseCase = usecases.NewLogoutUseCase(u.AuthUseCase, repos.BlacklistedTokenRepository)

	u.ChangeUserPasswordUseCase = usecases.NewChangeUserPasswordUseCase(u.AuthUseCase, repos.UserRepository)

	u.ChangeUserUsernameUseCase = usecases.NewChangeUserUsernameUseCase(u.AuthUseCase, repos.UserRepository)

	u.UserUseCase = usecases.NewUserUseCase(u.AuthUseCase, repos.UserRepository)

	u.BlacklistedTokenUseCase = usecases.NewBlacklistedTokenUseCase(repos.BlacklistedTokenRepository)

	return u
}

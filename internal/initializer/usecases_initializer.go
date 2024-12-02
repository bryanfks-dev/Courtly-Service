package initializer

import "main/domain/usecases"

// UseCases is a struct that holds all the use cases.
type UseCases struct {
	AuthUseCase             *usecases.AuthUseCase
	VerifyPasswordUseCase   *usecases.VerifyPasswordUseCase
	RegisterUseCase         *usecases.RegisterUseCase
	LoginUseCase            *usecases.LoginUseCase
	LogoutUseCase           *usecases.LogoutUseCase
	UserUseCase             *usecases.UserUseCase
	BlacklistedTokenUseCase *usecases.BlacklistedTokenUseCase
	VendorUseCase           *usecases.VendorUseCase
	CourtUseCase            *usecases.CourtUseCase
	ReviewUseCase           *usecases.ReviewUseCase
	BookingUseCase          *usecases.BookingUseCase
}

// InitUseCases is a function that initializes all the use cases.
//
// repos: The repositories.
//
// Returns a pointer to the UseCases struct.
func InitUseCases(repos *Repositories) *UseCases {
	u := &UseCases{}

	u.AuthUseCase = usecases.NewAuthUseCase()

	u.VerifyPasswordUseCase = usecases.NewVerifyPasswordUseCase(u.AuthUseCase, repos.UserRepository, repos.VendorRepository)

	u.RegisterUseCase = usecases.NewRegisterUseCase(u.AuthUseCase, repos.UserRepository)

	u.LoginUseCase = usecases.NewLoginUseCase(u.AuthUseCase, repos.UserRepository, repos.VendorRepository)

	u.LogoutUseCase = usecases.NewLogoutUseCase(u.AuthUseCase, repos.BlacklistedTokenRepository)

	u.UserUseCase = usecases.NewUserUseCase(u.AuthUseCase, repos.UserRepository)

	u.BlacklistedTokenUseCase = usecases.NewBlacklistedTokenUseCase(repos.BlacklistedTokenRepository)

	u.VendorUseCase = usecases.NewVendorUseCase(u.AuthUseCase, repos.VendorRepository)

	u.CourtUseCase = usecases.NewCourtUseCase(u.AuthUseCase, repos.CourtRepository)

	u.ReviewUseCase = usecases.NewReviewUseCase(u.AuthUseCase, repos.ReviewRepository)

	u.BookingUseCase = usecases.NewBookingUseCase(u.AuthUseCase, repos.BookingRepository)

	return u
}

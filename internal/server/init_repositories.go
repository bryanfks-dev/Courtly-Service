package server

import "main/internal/repository"

var (
	// userRepository is a variable that holds the user repository
	userRepository *repository.UserRepository

	// blacklistedTokenRepository is a variable that holds the blacklisted token repository
	blacklistedTokenRepository *repository.BlacklistedTokenRepository
)

// initRepositories is a function that initializes the repositories
//
// Returns void
func initRepositories() {
	userRepository = repository.NewUserRepository()

	blacklistedTokenRepository = repository.NewBlacklistedTokenRepository()
}

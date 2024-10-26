package initializer

import "main/internal/repository"

// Repositories is a struct that holds all the repositories.
type Repositories struct {
	UserRepository             *repository.UserRepository
	BlacklistedTokenRepository *repository.BlacklistedTokenRepository
}

// InitRepositories is a function that initializes all the repositories.
//
// Returns a pointer to the Repositories struct.
func InitRepositories() *Repositories {
	return &Repositories{
		UserRepository:             repository.NewUserRepository(),
		BlacklistedTokenRepository: repository.NewBlacklistedTokenRepository(),
	}
}

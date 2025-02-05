package usecases

import (
	"main/data/models"
	"main/internal/repository"
)

// BlacklistedTokenUseCase is a struct that defines the blacklisted token use case.
type BlacklistedTokenUseCase struct {
	BlacklistedTokenRepository *repository.BlacklistedTokenRepository
}

// NewBlacklistedTokenUseCase is a factory function that returns a new instance of the BlacklistedTokenUseCase.
//
// b: The blacklisted token repository.
//
// Returns a new instance of the BlacklistedTokenUseCase.
func NewBlacklistedTokenUseCase(b *repository.BlacklistedTokenRepository) *BlacklistedTokenUseCase {
	return &BlacklistedTokenUseCase{BlacklistedTokenRepository: b}
}

// AddBlacklistToken is a function that adds a token to the blacklist.
//
// token: The token to be added to the blacklist.
//
// Returns an error if the operation was not successful.
func (b *BlacklistedTokenUseCase) AddBlacklistToken(token *models.BlacklistedToken) error {
	// Add the token to the blacklist
	err := b.BlacklistedTokenRepository.Create(token)

	// If there was an error adding the token to the blacklist, log the error and return the error
	if err != nil {
		return err
	}

	return nil
}

// IsBlacklistedToken is a function that checks if a token is blacklisted.
//
// token: The token to be checked.
//
// Returns a boolean value.
func (b *BlacklistedTokenUseCase) IsBlacklistedToken(token string) bool {
	// Check if the token is blacklisted
	blacklisted, err := b.BlacklistedTokenRepository.IsBlacklisted(token)

	// If there was an error checking if the token is blacklisted, log the error and return false
	if err != nil {
		return false
	}

	return blacklisted
}

// ClearBlacklistToken is a function that deletes all the expired tokens from the blacklist.
//
// Returns an error if the operation was not successful.
func (b *BlacklistedTokenUseCase) ClearBlacklistToken() error {
	// Clear the blacklist token
	err := b.BlacklistedTokenRepository.Clear()

	// If there was an error clearing the blacklist token, log the error and return the error
	if err != nil {
		return err
	}

	return nil
}

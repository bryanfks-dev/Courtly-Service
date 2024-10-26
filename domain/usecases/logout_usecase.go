package usecases

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"
	"main/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

// LogoutUseCase is a struct that defines the logout use case.
type LogoutUseCase struct {
	blacklistedTokenRepository *repository.BlacklistedTokenRepository
	authUseCase                *AuthUseCase
}

// NewLogoutUseCase is a factory function that returns a new instance of the LogoutUseCase.
//
// b: The blacklisted token repository.
//
// Returns a new instance of the LogoutUseCase.
func NewLogoutUseCase(b *repository.BlacklistedTokenRepository, a *AuthUseCase) *LogoutUseCase {
	return &LogoutUseCase{blacklistedTokenRepository: b, authUseCase: a}
}

// BlacklistToken is a usecase that invalidates a token
// by removing it from the blacklisted_tokens table
//
// token: the token to invalidate
//
// Returns an error if the token could not be invalidated
func (l *LogoutUseCase) BlacklistToken(token *jwt.Token) error {
	// Get the token claims
	claims := l.authUseCase.DecodeToken(token)

	// Insert the token into the blacklisted_tokens table
	err := mysql.Conn.Create(&models.BlacklistedToken{
		Token:     token.Raw,
		ExpiresAt: claims.ExpiresAt.Time,
	}).Error

	// Return an error if any
	if err != nil {
		log.Println("Error blacklisting token: ", err)

		return err
	}

	return nil
}

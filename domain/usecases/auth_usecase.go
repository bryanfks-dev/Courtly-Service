package usecases

import (
	"errors"
	"log"
	"main/core/config"
	"main/core/enums"
	"main/domain/entities"
	"main/pkg/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// AuthUseCase is a struct that defines the usecase for authentication
type AuthUseCase struct{}

// NewAuthUseCase is a factory function that returns a new instance of the AuthUseCase struct
//
// Returns a new instance of the AuthUseCase struct
func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

// GenerateToken is a function that generates a JWT token.
//
// id: the id of the client
// clientType: the client type
//
// Returns a string containing the token and an error if there is any
func (a *AuthUseCase) GenerateToken(id uint, clientType enums.ClientType) (string, error) {
	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.JWTClaims{
		Id:         id,
		ClientType: clientType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(config.JWTConfig.Secret))

	// Check if there is an error
	if err != nil {
		log.Fatal("Error generating token: ", err)

		return "", err
	}

	return tokenString, nil
}

// ExtractToken is a usecase that extracts the token from the request header
//
// c: the echo.Context object
//
// Returns the token string and an error if there is any
func (a *AuthUseCase) ExtractToken(c echo.Context) (string, error) {
	// Extract the token from the request header
	tokenString := c.Request().Header.Get("Authorization")

	// Check if the token is blank
	if utils.IsBlank(tokenString) {
		return "", errors.New("token is missing")
	}

	// Sanitize the tokenString
	tokenString = strings.TrimSpace(tokenString)

	token := strings.Split(tokenString, " ")

	// Check if the token is in the correct format
	if token[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	return token[1], nil
}

// DecodeToken is a function that decodes a JWT token
//
// token: the token to decode
//
// Returns the decoded token.
func (a *AuthUseCase) DecodeToken(token *jwt.Token) *entities.JWTClaims {
	return token.Claims.(*entities.JWTClaims)
}

// HashPassword is a function that hashes a password.
//
// password: the password to hash
//
// Returns the hashed password and an error if there is one
func (a *AuthUseCase) HashPassword(password string) (string, error) {
	// Generate a hashed password from the password string
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// VerifyPassword is a function that verifies the password.
//
// password: The password to verify
// hashedPwd: The hashed password to compare with
//
// Retuns true if the password is correct, false otherwise
func (a *AuthUseCase) VerifyPassword(password string, hashedPwd string) bool {
	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))

	return err == nil
}

// VerifyToken is a function that verifies a JWT token.
//
// tokenString: the token to verify
//
// Returns a boolean indicating if the token is valid and a jwt.Token object
func (a *AuthUseCase) VerifyToken(tokenString string) (*jwt.Token, bool) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &entities.JWTClaims{}, func(t *jwt.Token) (any, error) {
		// Check if the token is signed with the correct signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.JWTConfig.Secret), nil
	})

	// Check if there is an error parsing token 
	// or the token is invalid
	if err != nil || !token.Valid {
		return nil, false
	}

	return token, true
}

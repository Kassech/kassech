package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

// Initialize JWT Secret when the app starts.
func InitJWTSecret() error {
	secret := os.Getenv("JWTToken")
	if secret == "" {
		return errors.New("JWT secret is not set in environment variables")
	}
	jwtSecret = secret
	return nil
}

// GenerateToken creates a JWT token for the given userID
func GenerateToken(userID uint) (string, error) {
	// Ensure that the JWT secret is initialized
	if jwtSecret == "" {
		return "", errors.New("JWT secret not initialized")
	}

	// Create the claims for the token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	// Generate the token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the preloaded JWT secret
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

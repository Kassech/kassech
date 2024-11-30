package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret string
var AccessTokenExpiration time.Duration
var RefreshTokenExpiration time.Duration

// Initialize JWT Secret and expiration times when the app starts.
func InitJWTSecret() error {
	// Get JWT secret from environment variables
	secret := os.Getenv("JWTToken")
	if secret == "" {
		return errors.New("JWT secret is not set in environment variables")
	}
	JwtSecret = secret

	// Get expiration times from environment variables
	accessTokenExp := os.Getenv("ACCESS_TOKEN_EXPIRATION") // e.g., "72h" or "2592000" (seconds)
	if accessTokenExp == "" {
		return errors.New("ACCESS_TOKEN_EXPIRATION is not set in environment variables")
	}

	// Parse access token expiration time
	parsedAccessTokenExp, err := time.ParseDuration(accessTokenExp)
	if err != nil {
		return errors.New("invalid format for ACCESS_TOKEN_EXPIRATION")
	}
	AccessTokenExpiration = parsedAccessTokenExp

	refreshTokenExp := os.Getenv("REFRESH_TOKEN_EXPIRATION") // e.g., "720h" or "2592000" (seconds)
	if refreshTokenExp == "" {
		return errors.New("REFRESH_TOKEN_EXPIRATION is not set in environment variables")
	}

	// Parse refresh token expiration time
	parsedRefreshTokenExp, err := time.ParseDuration(refreshTokenExp)
	if err != nil {
		return errors.New("invalid format for REFRESH_TOKEN_EXPIRATION")
	}
	RefreshTokenExpiration = parsedRefreshTokenExp

	return nil
}

// GenerateToken creates an access JWT token and a refresh JWT token for the given userID
func GenerateToken(userID uint) (string, string, error) {
	// Ensure that the JWT secret and expiration times are initialized
	if JwtSecret == "" || AccessTokenExpiration == 0 || RefreshTokenExpiration == 0 {
		return "", "", errors.New("JWT secret or expiration times not initialized")
	}

	// Create the claims for the access token (short-lived)
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(AccessTokenExpiration).Unix(), // Token expires in the duration specified in ACCESS_TOKEN_EXPIRATION
	}

	// Generate the access token with claims and signing method
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Sign the access token using the preloaded JWT secret
	accessTokenString, err := accessToken.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", "", err
	}

	// Create the claims for the refresh token (longer-lived)
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(RefreshTokenExpiration).Unix(), // Refresh token expires in the duration specified in REFRESH_TOKEN_EXPIRATION
	}

	// Generate the refresh token with claims and signing method
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Sign the refresh token using the preloaded JWT secret
	refreshTokenString, err := refreshToken.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", "", err
	}

	// Return both tokens
	return accessTokenString, refreshTokenString, nil
}

// RefreshTokenService is used to refresh an expired access token using a valid refresh token
func RefreshTokenService(refreshToken string) (string, error) {
	// Ensure that the JWT secret is initialized
	if JwtSecret == "" {
		return "", errors.New("JWT secret not initialized")
	}

	// Parse and validate the refresh token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	// Check for errors in parsing or invalid expiration
	if err != nil || claims["exp"] == nil || time.Now().Unix() > int64(claims["exp"].(float64)) {
		return "", errors.New("invalid or expired refresh token")
	}

	// Extract the user ID from the refresh token
	userID := uint(claims["user_id"].(float64))

	// Generate a new access token
	accessToken, _, err := GenerateToken(userID)
	if err != nil {
		return "", err
	}

	// Return the newly generated access token
	return accessToken, nil
}

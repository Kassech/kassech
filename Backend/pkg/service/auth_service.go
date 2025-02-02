package service

import (
	"context"
	"errors"
	"fmt"
	"kassech/backend/pkg/config"
	"kassech/backend/pkg/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret string
var AccessTokenExpiration time.Duration
var RefreshTokenExpiration time.Duration

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

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
	if err != nil || parsedAccessTokenExp <= 0 {
		return errors.New("invalid format or value for ACCESS_TOKEN_EXPIRATION")
	}
	AccessTokenExpiration = parsedAccessTokenExp

	refreshTokenExp := os.Getenv("REFRESH_TOKEN_EXPIRATION") // e.g., "720h" or "2592000" (seconds)
	if refreshTokenExp == "" {
		return errors.New("REFRESH_TOKEN_EXPIRATION is not set in environment variables")
	}

	// Parse refresh token expiration time
	parsedRefreshTokenExp, err := time.ParseDuration(refreshTokenExp)
	if err != nil || parsedRefreshTokenExp <= 0 {
		return errors.New("invalid format or value for REFRESH_TOKEN_EXPIRATION")
	}
	RefreshTokenExpiration = parsedRefreshTokenExp

	return nil
}

// GenerateToken creates an access JWT token and a refresh JWT token for the given userID
func GenerateToken(userID uint, role []string) (string, string, error) {
	// Ensure that the JWT secret and expiration times are initialized
	if JwtSecret == "" || AccessTokenExpiration == 0 || RefreshTokenExpiration == 0 {
		return "", "", errors.New("JWT secret or expiration times not initialized")
	}

	// Create the claims for the access token (short-lived)
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(AccessTokenExpiration).Unix(), // Token expires in the duration specified in ACCESS_TOKEN_EXPIRATION
		"role":    role,
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
func RefreshTokenService(refreshToken string) (string, string, error) {
	// Ensure that the JWT secret is initialized
	if JwtSecret == "" {
		return "", "", errors.New("JWT secret not initialized")
	}

	// Parse and validate the refresh token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})

	// Check for errors in parsing or invalid expiration
	if err != nil || claims["exp"] == nil || time.Now().Unix() > int64(claims["exp"].(float64)) {
		return "", "", errors.New("invalid or expired refresh token")
	}

	// Extract the user ID from the refresh token
	userID := uint(claims["user_id"].(float64))
	role, ok := claims["role"].([]interface{})
	if !ok {
		return "", "", errors.New("invalid role in refresh token")
	}
	roleStr := make([]string, len(role))
	for i, r := range role {
		roleStr[i] = fmt.Sprint(r)
	}
	// Redis key
	ctx := context.Background()
	redisKey := "refresh_token:" + fmt.Sprintf("%d", int(userID))

	// Check token in Redis
	val, redisErr := config.RedisClient.Get(ctx, redisKey).Result()
	if redisErr == nil && val == "active" {
		// Token is valid in Redis, proceed to generate a new access token
		accessToken, _, err := GenerateToken(userID, roleStr)
		fmt.Println("accessToken, _, err:", accessToken, err)
		if err != nil {
			return "", "", errors.New("failed to cache refresh token in Redis")
		}
		return accessToken, fmt.Sprintf("%d", int(userID)), nil

	}

	// If not in Redis, validate the token in the database
	var isValid bool
	err = database.DB.Raw(`
	SELECT COUNT(*) > 0
	FROM user_sessions
	WHERE token = ? AND is_active = TRUE AND expires_at > NOW()
`, refreshToken).Scan(&isValid).Error
	if err != nil || !isValid {
		return "", "", errors.New("refresh token is invalid or expired")
	}

	// Cache token in Redis for future requests
	err = config.RedisClient.Set(ctx, redisKey, "active", 24*time.Hour).Err()
	if err != nil {
		return "", "", errors.New("failed to cache refresh token in Redis")
	}

	// Generate a new access token
	accessToken, _, err := GenerateToken(userID, roleStr)
	if err != nil {
		return "", "", err
	}

	// Return the newly generated access token
	return accessToken, fmt.Sprintf("%d", int(userID)), nil
}

// ValidateToken validates the given JWT token and returns the user information if valid
func VerifyAccessToken(tokenString string) (*Claims, error) {
	// Ensure that the JWT secret is initialized
	if JwtSecret == "" {
		return nil, errors.New("JWT secret not initialized")
	}

	// Parse the token and validate the claims
	claims := &Claims{}
	// Parse the token with the claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method is HMAC and validate the signing key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JwtSecret), nil
	})

	// Check if there was an error parsing the token or if the token is invalid
	if err != nil {
		return nil, err
	}

	// Check if the token is valid (not expired)
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Check if the token has expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	// Return the claims (user info)
	return claims, nil
}

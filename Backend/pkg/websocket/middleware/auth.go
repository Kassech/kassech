package middleware

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type WebSocketAuth struct {
	jwtSecret string
}

func NewWebSocketAuth(jwtSecret string) *WebSocketAuth {
	return &WebSocketAuth{jwtSecret: jwtSecret}
}

func (a *WebSocketAuth) Authenticate(r *http.Request) (uint, error) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		return 0, errors.New("authorization token required")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token")
	}

	return uint(userID), nil
}

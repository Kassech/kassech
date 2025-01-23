package middleware

import (
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or invalid"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(service.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": token})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		userIDFloat64, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token claims"})
			c.Abort()
			return
		}
		userID := strconv.FormatFloat(userIDFloat64, 'f', -1, 64)
		role, ok := claims["role"].([]interface{})
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role in token claims"})
			c.Abort()
			return
		}
		var roles []string
		for _, r := range role {
			if roleStr, valid := r.(string); valid {
				roles = append(roles, roleStr)
			}
		}
		isAdmin := utils.Contains(roles, "Admin")
		c.Set("isAdmin", isAdmin)
		c.Set("userID", userID)
		c.Set("role", roles)
		c.Next()
	}
}

package controller

import (
	"fmt"
	"kassech/backend/pkg/database"
	"kassech/backend/pkg/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserSessionController struct {
	Service *service.UserSessionsService
}

func (c *UserSessionController) InvalidateToken(ctx *gin.Context) {
	token := ctx.Param("token")
	if err := c.Service.Logout(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate token"})
		return
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.JwtSecret), nil
	})

	// Check for errors in parsing or invalid expiration
	if err != nil || claims["exp"] == nil || time.Now().Unix() > int64(claims["exp"].(float64)) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid or expired refresh token"})
		return

	}

	// Extract the user ID from the refresh token
	userID := uint(claims["user_id"].(float64))

	redisKey := "refresh_token:" + fmt.Sprintf("%d", int(userID))
	if err := database.REDIS.Del(ctx, redisKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh Redis key"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Token invalidated successfully"})
}

func (c *UserSessionController) InvalidateAllSessions(ctx *gin.Context) {
	UserID := ctx.Param("id")

	if err := c.Service.DisableUser(UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate sessions"})
		return
	}

	// Remove all Redis keys related to user sessions
	redisKeyPattern := "refresh_token:%d*" + UserID
	keys, err := database.REDIS.Keys(ctx, redisKeyPattern).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Redis keys"})
		return
	}

	for _, key := range keys {
		if err := database.REDIS.Del(ctx, key).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove Redis keys"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All sessions invalidated successfully"})
}

package middleware

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		log.Println("userID:", userID) // Debug log

		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		var userRoles []models.UserRole
		database.DB.Where("user_id = ?", userID).Find(&userRoles)

		for _, userRole := range userRoles {
			var role models.Role
			database.DB.First(&role, userRole.RoleID)
			log.Println("ROLE:", string(role.RoleName))        // Debug log
			log.Println("requiredRole:", string(requiredRole)) // Debug log

			if role.RoleName == requiredRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		c.Abort()
	}
}

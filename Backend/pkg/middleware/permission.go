package middleware

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check if a user has a specific permission
func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
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

			var permissions []models.Permission
			database.DB.Model(&role).Association("Permissions").Find(&permissions)

			for _, permission := range permissions {
				if permission.PermissionName == requiredPermission {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}

// Check if a user has one or more permissions from a list
func AnyPermissionMiddleware(requiredPermissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
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

			var permissions []models.Permission
			database.DB.Model(&role).Association("Permissions").Find(&permissions)

			for _, permission := range permissions {
				for _, requiredPermission := range requiredPermissions {
					if permission.PermissionName == requiredPermission {
						c.Next()
						return
					}
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}

// Check if a user has all permissions from a list
func AllPermissionsMiddleware(requiredPermissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
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

			var permissions []models.Permission
			database.DB.Model(&role).Association("Permissions").Find(&permissions)

			permissionMap := make(map[string]bool)
			for _, permission := range permissions {
				permissionMap[permission.PermissionName] = true
			}

			for _, requiredPermission := range requiredPermissions {
				if !permissionMap[requiredPermission] {
					c.JSON(http.StatusForbidden, gin.H{"error": "Missing required permissions"})
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}

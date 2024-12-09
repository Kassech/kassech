package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRoleController struct {
	Service *service.UserRoleService
}

// CreateUserRole handles creating a new user role
func (ctrl *UserRoleController) CreateUserRole(c *gin.Context) {
	var userRole models.UserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.Create(&userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userRole)
}

// GetUserRole retrieves a user role by ID
func (ctrl *UserRoleController) GetUserRole(c *gin.Context) {
	id := c.Param("id")
	userRole, err := ctrl.Service.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User role not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRole)
}

// UpdateUserRole updates an existing user role
func (ctrl *UserRoleController) UpdateUserRole(c *gin.Context) {
	id := c.Param("id")
	var userRole models.UserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.Update(id, &userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRole)
}

// DeleteUserRole deletes a user role by ID
func (ctrl *UserRoleController) DeleteUserRole(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.Service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User role not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

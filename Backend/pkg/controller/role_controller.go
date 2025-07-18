package controller

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"kassech/backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	Service *service.RoleService
}

// CreateRole handles HTTP requests to create a new role
func (rc *RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRole, err := rc.Service.CreateRole(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdRole)
}

// UpdateRole handles HTTP requests to update an existing role
func (rc *RoleController) UpdateRole(c *gin.Context) {
	var role models.Role
	roleID := c.Param("id")

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roleIDUint, err := utils.StringToUint(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	updatedRole, err := rc.Service.UpdateRole(&role, roleIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRole)
}

// DeleteRole handles HTTP requests to delete a role
func (rc *RoleController) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	roleIDUint, err := utils.StringToUint(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	deletedRole, err := rc.Service.DeleteRoleByID(roleIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedRole)
}

// FindRoleByID handles HTTP requests to get a role by ID
func (rc *RoleController) FindRoleByID(c *gin.Context) {
	roleID := c.Param("id")
	roleIDUint, err := utils.StringToUint(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	role, err := rc.Service.FindRoleByID(roleIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// GetAllRole handles HTTP requests to get a role by ID
func (rc *RoleController) GetAllRoles(c *gin.Context) {

	role, err := rc.Service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

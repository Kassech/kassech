package controller

import (
	"net/http"
	"strconv"

	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
)

// PermissionController defines the controller for handling permission-related requests
type PermissionController struct {
	Service *service.PermissionService
}

// NewPermissionController creates a new PermissionController instance
func NewPermissionController(Service *service.PermissionService) *PermissionController {
	return &PermissionController{
		Service: Service,
	}
}

// CreatePermission handles the creation of a new permission
func (pc *PermissionController) CreatePermission(c *gin.Context) {
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdPermission, err := pc.Service.CreatePermission(&permission)
	if err != nil {
		log.Println("Error creating permission:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create permission"})
		return
	}

	c.JSON(http.StatusCreated, createdPermission)
}

// GetPermissionByID retrieves a permission by ID
func (pc *PermissionController) GetPermissionByID(c *gin.Context) {
	permissionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	permission, err := pc.Service.GetPermissionByID(uint(permissionID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.JSON(http.StatusOK, permission)
}

// GetAllPermissions retrieves all permissions
func (pc *PermissionController) GetAllPermissions(c *gin.Context) {
	permissions, err := pc.Service.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve permissions"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

// UpdatePermission updates an existing permission
func (pc *PermissionController) UpdatePermission(c *gin.Context) {
	permissionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	permission.ID = uint(permissionID)

	updatedPermission, err := pc.Service.UpdatePermission(&permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update permission"})
		return
	}

	c.JSON(http.StatusOK, updatedPermission)
}

// DeletePermission deletes a permission by ID
func (pc *PermissionController) DeletePermission(c *gin.Context) {
	permissionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	deletedPermission, err := pc.Service.DeletePermission(uint(permissionID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.JSON(http.StatusOK, deletedPermission)
}

// AttachRoleToPermission attaches a role to a permission
func (pc *PermissionController) AttachRoleToPermission(c *gin.Context) {
	permissionID, err := strconv.Atoi(c.Param("permission_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Permission ID"})
		return
	}

	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Role ID"})
		return
	}

	err = pc.Service.AttachRoleToPermission(uint(permissionID), uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to attach role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role attached successfully"})
}

// DetachRoleFromPermission detaches a role from a permission
func (pc *PermissionController) DetachRoleFromPermission(c *gin.Context) {
	permissionID, err := strconv.Atoi(c.Param("permission_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Permission ID"})
		return
	}

	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Role ID"})
		return
	}

	err = pc.Service.DetachRoleFromPermission(uint(permissionID), uint(roleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to detach role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role detached successfully"})
}

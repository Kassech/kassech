package controller

import (
	"net/http"
	"strconv"

	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/service"

	"github.com/gin-gonic/gin"
)

type RolePermissionController struct {
	service *service.RolePermissionService
}

func NewRolePermissionController(service *service.RolePermissionService) *RolePermissionController {
	return &RolePermissionController{
		service: service,
	}
}

// CreateRolePermission handles the creation of a RolePermission
func (c *RolePermissionController) CreateRolePermission(ctx *gin.Context) {
	var rolePermission models.RolePermission

	// Bind JSON body to RolePermission struct
	if err := ctx.ShouldBindJSON(&rolePermission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the service to create the role permission
	createdRolePermission, err := c.service.Create(&rolePermission)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created role permission
	ctx.JSON(http.StatusCreated, createdRolePermission)
}

// GetRolePermissionByID handles fetching a RolePermission by ID
func (c *RolePermissionController) GetRolePermissionByID(ctx *gin.Context) {
	rolePermissionID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service to fetch the role permission
	rolePermission, err := c.service.GetByID(uint(rolePermissionID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "RolePermission not found"})
		return
	}

	// Return the found role permission
	ctx.JSON(http.StatusOK, rolePermission)
}

// GetAllRolePermissions handles fetching all RolePermissions
func (c *RolePermissionController) GetAllRolePermissions(ctx *gin.Context) {
	// Call the service to fetch all role permissions
	rolePermissions, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of role permissions
	ctx.JSON(http.StatusOK, rolePermissions)
}

// DeleteRolePermission handles deleting a RolePermission by ID
func (c *RolePermissionController) DeleteRolePermission(ctx *gin.Context) {
	rolePermissionID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service to delete the role permission
	deletedRolePermission, err := c.service.DeleteByID(uint(rolePermissionID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "RolePermission not found"})
		return
	}

	// Return the deleted role permission
	ctx.JSON(http.StatusOK, deletedRolePermission)
}

// DeleteRolePermissionByRoleAndPermission handles deleting a RolePermission by RoleID and PermissionID
func (c *RolePermissionController) DeleteRolePermissionByRoleAndPermission(ctx *gin.Context) {
	roleID, err := strconv.Atoi(ctx.Param("role_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Role ID"})
		return
	}

	permissionID, err := strconv.Atoi(ctx.Param("permission_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Permission ID"})
		return
	}

	// Call the service to delete the role permission by role and permission
	if err := c.service.DeleteByRoleAndPermission(uint(roleID), uint(permissionID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "RolePermission deleted successfully"})
}

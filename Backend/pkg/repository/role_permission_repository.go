package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type RolePermissionRepository struct{}

// Create a new RolePermission in the database
func (rpr *RolePermissionRepository) Create(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	if err := rolePermission.Validate(); err != nil {
		return nil, err
	}
	if err := database.DB.Create(rolePermission).Error; err != nil {
		return nil, err
	}
	return rolePermission, nil
}

// Find a RolePermission by ID
func (rpr *RolePermissionRepository) FindByID(rolePermissionID uint) (*models.RolePermission, error) {
	var rolePermission models.RolePermission
	if err := database.DB.First(&rolePermission.RoleID, rolePermissionID).Error; err != nil {
		return nil, err
	}
	return &rolePermission, nil
}

// GetAll RolePermissions
func (rpr *RolePermissionRepository) GetAll() (*[]models.RolePermission, error) {
	var rolePermissions []models.RolePermission
	if err := database.DB.Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return &rolePermissions, nil
}

// Delete a RolePermission by ID
func (rpr *RolePermissionRepository) DeleteByID(rolePermissionID uint) (*models.RolePermission, error) {
	var rolePermission models.RolePermission
	// Find the RolePermission first to ensure it exists
	if err := database.DB.First(&rolePermission, rolePermissionID).Error; err != nil {
		return nil, err
	}

	// Delete the RolePermission after fetching it
	if err := database.DB.Delete(&rolePermission).Error; err != nil {
		return nil, err
	}

	return &rolePermission, nil
}

// FindByRoleIDAndPermissionID finds a RolePermission by RoleID and PermissionID
func (rpr *RolePermissionRepository) FindByRoleIDAndPermissionID(roleID, permissionID uint) (*models.RolePermission, error) {
	var rolePermission models.RolePermission
	if err := database.DB.Where("role_id = ? AND permission_id = ?", roleID, permissionID).First(&rolePermission).Error; err != nil {
		return nil, err
	}
	return &rolePermission, nil
}

// DeleteByRoleIDAndPermissionID deletes a RolePermission by RoleID and PermissionID
func (rpr *RolePermissionRepository) DeleteByRoleIDAndPermissionID(roleID, permissionID uint) error {
	if err := database.DB.Unscoped().Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&models.RolePermission{}).Error; err != nil {
		return err
	}
	return nil
}

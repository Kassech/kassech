package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

// PermissionRepository handles operations related to permissions in the database
type PermissionRepository struct{}

// Create a new permission in the database
func (pr *PermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	// Validate the permission before creating
	if err := permission.Validate(); err != nil {
		return nil, err
	}

	if err := database.DB.Create(permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

// Find a permission by ID
func (pr *PermissionRepository) FindByID(permissionID uint) (*models.Permission, error) {
	var permission models.Permission
	if err := database.DB.First(&permission, permissionID).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetAll permissions from the database
func (pr *PermissionRepository) GetAll() (*[]models.Permission, error) {
	var permissions []models.Permission
	if err := database.DB.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return &permissions, nil
}

// Update an existing permission by ID
func (pr *PermissionRepository) Update(permission *models.Permission) (*models.Permission, error) {
	// Check if the permission exists
	existingPermission, err := pr.FindByID(permission.ID)
	if err != nil {
		return nil, err
	}

	// Validate the updated permission
	if err := permission.Validate(); err != nil {
		return nil, err
	}

	// Use GORM's Updates method to update only the fields provided in the input
	if err := database.DB.Model(existingPermission).Updates(permission).Error; err != nil {
		return nil, err
	}

	return existingPermission, nil
}

// Delete a permission by ID
func (pr *PermissionRepository) DeleteByID(permissionID uint) (*models.Permission, error) {
	var permission models.Permission
	// Find the permission first to ensure it exists
	if err := database.DB.First(&permission, permissionID).Error; err != nil {
		return nil, err
	}

	// Delete the permission after fetching it
	if err := database.DB.Delete(&permission).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

// Attach a role to a permission (many-to-many relationship)
func (pr *PermissionRepository) AttachRoleToPermission(permissionID uint, roleID uint) error {
	// Find permission and role to ensure they exist
	permission, err := pr.FindByID(permissionID)
	if err != nil {
		return err
	}

	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return err
	}

	// Attach the role to the permission
	if err := database.DB.Model(permission).Association("Roles").Append(&role); err != nil {
		return err
	}

	return nil
}

// Detach a role from a permission
func (pr *PermissionRepository) DetachRoleFromPermission(permissionID uint, roleID uint) error {
	// Find permission and role to ensure they exist
	permission, err := pr.FindByID(permissionID)
	if err != nil {
		return err
	}

	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return err
	}

	// Detach the role from the permission
	if err := database.DB.Model(permission).Association("Roles").Delete(&role); err != nil {
		return err
	}

	return nil
}

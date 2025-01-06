package service

import (
	"errors"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type PermissionService struct {
	Repo *repository.PermissionRepository
}

// CreatePermission creates a new permission
func (ps *PermissionService) CreatePermission(permission *models.Permission) (*models.Permission, error) {
	// Use repository to create permission
	createdPermission, err := ps.Repo.Create(permission)
	if err != nil {
		return nil, err
	}
	return createdPermission, nil
}

// GetPermissionByID retrieves a permission by ID
func (ps *PermissionService) GetPermissionByID(permissionID uint) (*models.Permission, error) {
	permission, err := ps.Repo.FindByID(permissionID)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

// GetAllPermissions retrieves all permissions
func (ps *PermissionService) GetAllPermissions() (*[]models.Permission, error) {
	permissions, err := ps.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// UpdatePermission updates an existing permission
func (ps *PermissionService) UpdatePermission(permission *models.Permission) (*models.Permission, error) {
	// Check if permission exists first
	_, err := ps.Repo.FindByID(permission.ID)
	if err != nil {
		return nil, errors.New("permission not found")
	}

	// Update the permission
	updatedPermission, err := ps.Repo.Update(permission)
	if err != nil {
		return nil, err
	}
	return updatedPermission, nil
}

// DeletePermission deletes a permission by ID
func (ps *PermissionService) DeletePermission(permissionID uint) (*models.Permission, error) {
	// Check if permission exists first
	_, err := ps.Repo.FindByID(permissionID)
	if err != nil {
		return nil, errors.New("permission not found")
	}

	// Delete the permission
	deletedPermission, err := ps.Repo.DeleteByID(permissionID)
	if err != nil {
		return nil, err
	}
	return deletedPermission, nil
}

// AttachRoleToPermission attaches a role to a permission
func (ps *PermissionService) AttachRoleToPermission(permissionID uint, roleID uint) error {
	// Attach role to permission
	err := ps.Repo.AttachRoleToPermission(permissionID, roleID)
	if err != nil {
		return err
	}
	return nil
}

// DetachRoleFromPermission detaches a role from a permission
func (ps *PermissionService) DetachRoleFromPermission(permissionID uint, roleID uint) error {
	// Detach role from permission
	err := ps.Repo.DetachRoleFromPermission(permissionID, roleID)
	if err != nil {
		return err
	}
	return nil
}

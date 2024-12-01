package service

import (
	"errors"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type PermissionService struct {
	permissionRepo *repository.PermissionRepository
}

// NewPermissionService creates a new instance of PermissionService
func NewPermissionService(permissionRepo *repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		permissionRepo: permissionRepo,
	}
}

// CreatePermission creates a new permission
func (ps *PermissionService) CreatePermission(permission *models.Permission) (*models.Permission, error) {
	// Use repository to create permission
	createdPermission, err := ps.permissionRepo.Create(permission)
	if err != nil {
		return nil, err
	}
	return createdPermission, nil
}

// GetPermissionByID retrieves a permission by ID
func (ps *PermissionService) GetPermissionByID(permissionID uint) (*models.Permission, error) {
	permission, err := ps.permissionRepo.FindByID(permissionID)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

// GetAllPermissions retrieves all permissions
func (ps *PermissionService) GetAllPermissions() (*[]models.Permission, error) {
	permissions, err := ps.permissionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// UpdatePermission updates an existing permission
func (ps *PermissionService) UpdatePermission(permission *models.Permission) (*models.Permission, error) {
	// Check if permission exists first
	_, err := ps.permissionRepo.FindByID(permission.ID)
	if err != nil {
		return nil, errors.New("permission not found")
	}

	// Update the permission
	updatedPermission, err := ps.permissionRepo.Update(permission)
	if err != nil {
		return nil, err
	}
	return updatedPermission, nil
}

// DeletePermission deletes a permission by ID
func (ps *PermissionService) DeletePermission(permissionID uint) (*models.Permission, error) {
	// Check if permission exists first
	_, err := ps.permissionRepo.FindByID(permissionID)
	if err != nil {
		return nil, errors.New("permission not found")
	}

	// Delete the permission
	deletedPermission, err := ps.permissionRepo.DeleteByID(permissionID)
	if err != nil {
		return nil, err
	}
	return deletedPermission, nil
}

// AttachRoleToPermission attaches a role to a permission
func (ps *PermissionService) AttachRoleToPermission(permissionID uint, roleID uint) error {
	// Attach role to permission
	err := ps.permissionRepo.AttachRoleToPermission(permissionID, roleID)
	if err != nil {
		return err
	}
	return nil
}

// DetachRoleFromPermission detaches a role from a permission
func (ps *PermissionService) DetachRoleFromPermission(permissionID uint, roleID uint) error {
	// Detach role from permission
	err := ps.permissionRepo.DetachRoleFromPermission(permissionID, roleID)
	if err != nil {
		return err
	}
	return nil
}

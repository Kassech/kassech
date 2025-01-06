package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type RolePermissionService struct {
	Repo *repository.RolePermissionRepository
}

// Create a new RolePermission
func (s *RolePermissionService) Create(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	// Validate the RolePermission object
	if err := rolePermission.Validate(); err != nil {
		return nil, err
	}
	// Call the repository to create the RolePermission
	return s.Repo.Create(rolePermission)
}

// Get a RolePermission by ID
func (s *RolePermissionService) GetByID(rolePermissionID uint) (*models.RolePermission, error) {
	return s.Repo.FindByID(rolePermissionID)
}

// Get all RolePermissions
func (s *RolePermissionService) GetAll() (*[]models.RolePermission, error) {
	return s.Repo.GetAll()
}

// Delete a RolePermission by ID
func (s *RolePermissionService) DeleteByID(rolePermissionID uint) (*models.RolePermission, error) {
	return s.Repo.DeleteByID(rolePermissionID)
}

// Delete a RolePermission by RoleID and PermissionID
func (s *RolePermissionService) DeleteByRoleAndPermission(roleID, permissionID uint) error {
	return s.Repo.DeleteByRoleIDAndPermissionID(roleID, permissionID)
}

// Assign a RolePermission (optional, if you need some additional logic)
func (s *RolePermissionService) AssignRolePermission(roleID, permissionID uint) (*models.RolePermission, error) {
	// Check if the role permission already exists
	existingPermission, err := s.Repo.FindByRoleIDAndPermissionID(roleID, permissionID)
	if err != nil {
		// If not found, create a new role permission
		newRolePermission := &models.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		}
		return s.Repo.Create(newRolePermission)
	}
	return existingPermission, nil
}

package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type RolePermissionService struct {
	repo *repository.RolePermissionRepository
}

func NewRolePermissionService(repo *repository.RolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{
		repo: repo,
	}
}

// Create a new RolePermission
func (s *RolePermissionService) Create(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	// Validate the RolePermission object
	if err := rolePermission.Validate(); err != nil {
		return nil, err
	}

	// Call the repository to create the RolePermission
	return s.repo.Create(rolePermission)
}

// Get a RolePermission by ID
func (s *RolePermissionService) GetByID(rolePermissionID uint) (*models.RolePermission, error) {
	return s.repo.FindByID(rolePermissionID)
}

// Get all RolePermissions
func (s *RolePermissionService) GetAll() (*[]models.RolePermission, error) {
	return s.repo.GetAll()
}

// Delete a RolePermission by ID
func (s *RolePermissionService) DeleteByID(rolePermissionID uint) (*models.RolePermission, error) {
	return s.repo.DeleteByID(rolePermissionID)
}

// Delete a RolePermission by RoleID and PermissionID
func (s *RolePermissionService) DeleteByRoleAndPermission(roleID, permissionID uint) error {
	return s.repo.DeleteByRoleIDAndPermissionID(roleID, permissionID)
}

// Assign a RolePermission (optional, if you need some additional logic)
func (s *RolePermissionService) AssignRolePermission(roleID, permissionID uint) (*models.RolePermission, error) {
	// Check if the role permission already exists
	existingPermission, err := s.repo.FindByRoleIDAndPermissionID(roleID, permissionID)
	if err != nil {
		// If not found, create a new role permission
		newRolePermission := &models.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		}
		return s.repo.Create(newRolePermission)
	}
	return existingPermission, nil
}

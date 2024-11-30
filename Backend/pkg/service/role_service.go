package service

import (
	"errors"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type RoleService struct {
	Repo *repository.RoleRepository
}

// NewRoleService creates and returns a new instance of RoleService
func NewRoleService(Repo *repository.RoleRepository) *RoleService {
	return &RoleService{Repo: Repo}
}

// CreateRole creates a new role
func (rs *RoleService) CreateRole(role *models.Role) (*models.Role, error) {
	// Validate the role (optional: you could also perform some additional checks here)
	if err := role.Validate(); err != nil {
		return nil, err
	}

	// Create the role in the database
	createdRole, err := rs.Repo.Create(role)
	if err != nil {
		return nil, err
	}

	return createdRole, nil
}

// UpdateRole updates an existing role
func (rs *RoleService) UpdateRole(role *models.Role, roleID uint) (*models.Role, error) {
	// Validate the role (optional: you could also perform some additional checks here)
	if err := role.Validate(); err != nil {
		return nil, err
	}

	// Update the role in the database
	updatedRole, err := rs.Repo.Update(role, roleID)
	if err != nil {
		return nil, err
	}

	return updatedRole, nil
}

// FindRoleByID retrieves a role by its ID
func (rs *RoleService) FindRoleByID(roleID uint) (*models.Role, error) {
	// Retrieve the role from the repository
	role, err := rs.Repo.FindByID(roleID)
	if err != nil {
		return nil, err
	}

	return role, nil
}

// DeleteRoleByID deletes a role by its ID
func (rs *RoleService) DeleteRoleByID(roleID uint) (*models.Role, error) {
	// First, check if the role exists
	role, err := rs.Repo.FindByID(roleID)
	if err != nil {
		return nil, err
	}

	// Prevent deletion of system roles (optional logic)
	// For example, you could have a system role that should never be deleted
	if role.RoleName == "Admin" {
		return nil, errors.New("cannot delete system role 'Admin'")
	}

	// Proceed with the deletion
	deletedRole, err := rs.Repo.DeleteByID(roleID)
	if err != nil {
		return nil, err
	}

	return deletedRole, nil
}

// FindRoleByID retrieves a role by its ID
func (rs *RoleService) GetAllRoles() (*[]models.Role, error) {
	// Retrieve the role from the repository
	role, err := rs.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return role, nil
}

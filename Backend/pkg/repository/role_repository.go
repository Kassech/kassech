package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type RoleRepository struct{}

// Create a new role in the database
func (rr *RoleRepository) Create(role *models.Role) (*models.Role, error) {
	if err := database.DB.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

// Find a role by ID
func (rr *RoleRepository) FindByID(roleID uint) (*models.Role, error) {
	var role models.Role
	if err := database.DB.Preload("Permissions").First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// GetAll a role by ID
func (rr *RoleRepository) GetAll() (*[]models.Role, error) {
	var role []models.Role
	if err := database.DB.Preload("Permissions").Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Update an existing role by ID
func (rr *RoleRepository) Update(role *models.Role, roleID uint) (*models.Role, error) {
	// Check if the role exists
	existingRole, err := rr.FindByID(roleID)
	if err != nil {
		return nil, err
	}

	// Use GORM's Updates method to update only the fields provided in the input
	if err := database.DB.Model(existingRole).Updates(role).Error; err != nil {
		return nil, err
	}

	return existingRole, nil
}

// Delete a role by ID
func (rr *RoleRepository) DeleteByID(roleID uint) (*models.Role, error) {
	var role models.Role
	// Find the role first to ensure it exists
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return nil, err
	}

	// Delete the role after fetching it
	if err := database.DB.Delete(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

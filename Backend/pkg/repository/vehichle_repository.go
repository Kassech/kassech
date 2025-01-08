package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type VehicleRepository struct{}

// Create a new role in the database
func (rr *VehicleRepository) Create(role *models.Vehicle) (*models.Vehicle, error) {
	if err := database.DB.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

// Find a role by ID
func (rr *VehicleRepository) FindByID(roleID uint) (*models.Vehicle, error) {
	var role models.Vehicle
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// GetAll a role by ID
func (rr *VehicleRepository) GetAll() (*[]models.Vehicle, error) {
	var role []models.Vehicle
	if err := database.DB.Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Update an existing role by ID
func (rr *VehicleRepository) Update(role *models.Vehicle, roleID uint) (*models.Vehicle, error) {
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
func (rr *VehicleRepository) DeleteByID(roleID uint) (*models.Vehicle, error) {
	var role models.Vehicle
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

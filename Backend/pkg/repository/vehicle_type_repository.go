package repository

import (
	"fmt"
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type VehicleTypeRepository struct{}

// Create a new role in the database
func (rr *VehicleTypeRepository) Create(role *models.VehicleType) (*models.VehicleType, error) {
	if err := database.DB.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

// Find a role by ID
func (rr *VehicleTypeRepository) FindByID(roleID uint, forceDelete bool) (*models.VehicleType, error) {
	var role models.VehicleType
	if forceDelete {
		if err := database.DB.Unscoped().First(&role, roleID).Error; err != nil {
			return nil, err
		}
	} else {
		if err := database.DB.First(&role, roleID).Error; err != nil {
			return nil, err
		}
	}
	return &role, nil
}

// GetAll a role by ID
func (rr *VehicleTypeRepository) GetAll() (*[]models.VehicleType, error) {
	var role []models.VehicleType
	if err := database.DB.Find(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Update an existing role by ID
func (rr *VehicleTypeRepository) Update(role *models.VehicleType, roleID uint) (*models.VehicleType, error) {
	// Check if the role exists
	existingRole, err := rr.FindByID(roleID, false)
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
func (rr *VehicleTypeRepository) DeleteByID(roleID uint, forceDelete bool) (*models.VehicleType, error) {
	fmt.Println("forceDelete:", forceDelete)
	var role models.VehicleType
	// Optionally perform an unscoped delete
	var query = database.DB
	if forceDelete {
		query = query.Unscoped()
	}
	// Find the role first to ensure it exists
	if err := query.First(&role, roleID).Error; err != nil {
		return nil, err
	}

	// Delete the role after fetching it
	if err := query.Delete(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

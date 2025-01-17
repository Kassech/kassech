package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type VehicleRepository struct{}

// Create a new vehicle in the database
func (rr *VehicleRepository) Create(vehicle *models.Vehicle) (*models.Vehicle, error) {
	if err := database.DB.Create(vehicle).Error; err != nil {
		return nil, err
	}
	return vehicle, nil
}

// Find a vehicle by ID
func (rr *VehicleRepository) FindByID(roleID uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	if err := database.DB.Preload("Owner").Preload("Type").First(&vehicle, roleID).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetAll a vehicle by ID
func (rr *VehicleRepository) GetAll(page, perPage int, search, ownerID, typeID string) ([]models.Vehicle, int64, error) {
	var vehicles []models.Vehicle
	var total int64

	query := database.DB.Preload("Owner").Preload("Type").Model(&models.Vehicle{})

	if search != "" {
		query = query.Where("license_number ILIKE ? OR vin ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if ownerID != "" {
		query = query.Where("owner_id = ?", ownerID)
	}

	if typeID != "" {
		query = query.Where("type_id = ?", typeID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * perPage).Limit(perPage).Find(&vehicles).Error
	if err != nil {
		return nil, 0, err
	}

	return vehicles, total, nil
}

// Update an existing vehicle by ID
func (rr *VehicleRepository) Update(vehicle *models.Vehicle, roleID uint) (*models.Vehicle, error) {
	// Check if the vehicle exists
	existingRole, err := rr.FindByID(roleID)
	if err != nil {
		return nil, err
	}

	// Use GORM's Updates method to update only the fields provided in the input
	if err := database.DB.Model(existingRole).Updates(vehicle).Error; err != nil {
		return nil, err
	}

	return existingRole, nil
}

// Delete a vehicle by ID
func (rr *VehicleRepository) DeleteByID(roleID uint, forceDelete bool) (*models.Vehicle, error) {
	var vehicle models.Vehicle

	// Optionally perform an unscoped delete
	var query = database.DB
	if forceDelete {
		query = query.Unscoped()
	}
	// Find the vehicle first to ensure it exists
	if err := query.First(&vehicle, roleID).Error; err != nil {
		return nil, err
	}

	// Delete the vehicle after fetching it
	if err := query.Delete(&vehicle).Error; err != nil {
		return nil, err
	}

	return &vehicle, nil
}

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

// Find a vehicle by ID with active driver
func (rr *VehicleRepository) FindByID(vehicleID uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := database.DB.
		Preload("Owner").
		Preload("Type").
		Preload("Driver.User").
		First(&vehicle, vehicleID).Error

	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetAll vehicles with optional filters and active drivers
func (rr *VehicleRepository) GetAll(page, perPage int, search, ownerID, typeID string) ([]models.Vehicle, int64, error) {
	var vehicles []models.Vehicle
	var total int64

	query := database.DB.
		Preload("Owner").
		Preload("Type").
		Preload("Driver.User").
		Model(&models.Vehicle{})

	if search != "" {
		query = query.Where("license_number ILIKE ? OR vin ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if ownerID != "" {
		query = query.Where("owner_id = ?", ownerID)
	}

	if typeID != "" {
		query = query.Where("type_id = ?", typeID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset((page - 1) * perPage).Limit(perPage).Find(&vehicles).Error
	return vehicles, total, err
}

// Update an existing vehicle by ID
func (rr *VehicleRepository) Update(vehicle *models.Vehicle, vehicleID uint) (*models.Vehicle, error) {
	existingVehicle, err := rr.FindByID(vehicleID)
	if err != nil {
		return nil, err
	}

	if err := database.DB.Model(existingVehicle).Updates(vehicle).Error; err != nil {
		return nil, err
	}
	return existingVehicle, nil
}

// Delete a vehicle by ID
func (rr *VehicleRepository) DeleteByID(vehicleID uint, forceDelete bool) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	query := database.DB
	if forceDelete {
		query = query.Unscoped()
	}

	if err := query.First(&vehicle, vehicleID).Error; err != nil {
		return nil, err
	}

	if err := query.Delete(&vehicle).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

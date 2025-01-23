package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type LocationRepository struct{}

func (lr *LocationRepository) Create(location *models.VehicleGPSLog) (*models.VehicleGPSLog, error) {
	return location, database.DB.Create(location).Error
}

func (lr *LocationRepository) FindByID(locationID uint) (*models.VehicleGPSLog, error) {
	var location models.VehicleGPSLog
	if err := database.DB.First(&location, locationID).Error; err != nil {
		return nil, err
	}
	return &location, nil
}
func (lr *LocationRepository) GetAll(page, perPage int, search string, vehicleID uint, pathID uint) ([]models.VehicleGPSLog, int64, error) {
	var locations []models.VehicleGPSLog
	var total int64

	query := database.DB.Model(&models.VehicleGPSLog{})

	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if vehicleID != 0 {
		query = query.Where("vehicle_id = ?", vehicleID)
	}

	if pathID != 0 {
		query = query.Where("path_id = ?", pathID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * perPage).Limit(perPage).Find(&locations).Error
	if err != nil {
		return nil, 0, err
	}

	return locations, total, nil
}
func (lr *LocationRepository) DeleteByID(locationID uint, forceDelete bool) (*models.VehicleGPSLog, error) {
	var location models.VehicleGPSLog
	if err := database.DB.First(&location, locationID).Error; err != nil {
		return nil, err
	}
	if forceDelete {
		if err := database.DB.Unscoped().Delete(&location).Error; err != nil {
			return nil, err
		}
	} else {
		if err := database.DB.Delete(&location).Error; err != nil {
			return nil, err
		}
	}
	return &location, nil
}

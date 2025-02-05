package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type PathRepository struct{}

// Create a new path
func (pr *PathRepository) Create(path *models.Path) (*models.Path, error) {
	if err := database.DB.Create(path).Error; err != nil {
		return nil, err
	}
	return path, nil
}

// Find a path by ID
func (pr *PathRepository) FindByID(pathID uint) (*models.Path, error) {
	var path models.Path
	if err := database.DB.
		Preload("Route").
		Preload("Route.StationA").
		Preload("Route.StationB").
		First(&path, pathID).Error; err != nil {
		return nil, err
	}
	return &path, nil
}

// Find a path by Station ID
func (pr *PathRepository) FindPathsByStationID(stationID uint) ([]models.Path, error) {
	var paths []models.Path
	err := database.DB.
		Preload("Route.StationA").
		Preload("Route.StationB").
		Joins("JOIN routes ON paths.route_id = routes.id").
		Where("routes.location_a = ?", stationID).
		Find(&paths).
		Error

	if err != nil {
		return nil, err
	}
	return paths, nil
}

// Get all paths
func (pr *PathRepository) GetAll(page, perPage int, search string) ([]models.Path, int64, error) {
	var paths []models.Path
	var total int64

	query := database.DB.Model(&models.Path{}).Preload("Route").Preload("Route.StationA").Preload("Route.StationB")

	if search != "" {
		query = query.Where("path_name ILIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * perPage).Limit(perPage).Find(&paths).Error
	if err != nil {
		return nil, 0, err
	}

	return paths, total, nil
}

// Update an existing path by ID
func (pr *PathRepository) Update(path *models.Path, pathID uint) (*models.Path, error) {
	// Check if the path exists
	existingPath, err := pr.FindByID(pathID)
	if err != nil {
		return nil, err
	}

	// Use GORM's Updates method to update only the fields provided in the input
	if err := database.DB.Model(existingPath).Updates(path).Error; err != nil {
		return nil, err
	}

	return existingPath, nil
}

// Delete a path by ID
func (pr *PathRepository) DeleteByID(pathID uint) (*models.Path, error) {
	var path models.Path
	// Find the path first to ensure it exists
	if err := database.DB.First(&path, pathID).Error; err != nil {
		return nil, err
	}

	// Delete the path after fetching it
	if err := database.DB.Delete(&path).Error; err != nil {
		return nil, err
	}

	return &path, nil
}

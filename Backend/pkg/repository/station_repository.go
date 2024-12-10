package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type StationRepository struct{}

// Create a new station in the database
func (rr *StationRepository) Create(station *models.Station) (*models.Station, error) {
	if err := database.DB.Create(station).Error; err != nil {
		return nil, err
	}
	return station, nil
}

// Find a station by ID
func (rr *StationRepository) FindByID(stationID uint) (*models.Station, error) {
	var station models.Station
	if err := database.DB.First(&station, stationID).Error; err != nil {
		return nil, err
	}
	return &station, nil
}

// GetAll a station by ID
func (rr *StationRepository) GetAll() (*[]models.Station, error) {
	var station []models.Station
	if err := database.DB.Find(&station).Error; err != nil {
		return nil, err
	}
	return &station, nil
}

// Update an existing station by ID
func (rr *StationRepository) Update(station *models.Station, stationID uint) (*models.Station, error) {
	// Check if the station exists
	existingstation, err := rr.FindByID(stationID)
	if err != nil {
		return nil, err
	}

	// Use GORM's Updates method to update only the fields provided in the input
	if err := database.DB.Model(existingstation).Updates(station).Error; err != nil {
		return nil, err
	}

	return existingstation, nil
}

// Delete a station by ID
func (rr *StationRepository) DeleteByID(stationID uint) (*models.Station, error) {
	var station models.Station
	// Find the station first to ensure it exists
	if err := database.DB.First(&station, stationID).Error; err != nil {
		return nil, err
	}

	// Delete the station after fetching it
	if err := database.DB.Delete(&station).Error; err != nil {
		return nil, err
	}

	return &station, nil
}

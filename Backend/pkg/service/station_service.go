package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type StationService struct {
	Repo *repository.StationRepository
}

// CreateStation creates a new station
func (rs *StationService) CreateStation(station *models.Station) (*models.Station, error) {
	// Create the station in the database
	createdStation, err := rs.Repo.Create(station)
	if err != nil {
		return nil, err
	}

	return createdStation, nil
}

// UpdateStation updates an existing station
func (rs *StationService) UpdateStation(station *models.Station, roleID uint) (*models.Station, error) {
	// Update the station in the database
	updatedStation, err := rs.Repo.Update(station, roleID)
	if err != nil {
		return nil, err
	}

	return updatedStation, nil
}

// FindStationByID retrieves a station by its ID
func (rs *StationService) FindStationByID(roleID uint) (*models.Station, error) {
	// Retrieve the station from the repository
	station, err := rs.Repo.FindByID(roleID)
	if err != nil {
		return nil, err
	}

	return station, nil
}

// DeleteStationByID deletes a station by its ID
func (rs *StationService) DeleteStationByID(roleID uint) (*models.Station, error) {
	// Proceed with the deletion
	deletedStation, err := rs.Repo.DeleteByID(roleID)
	if err != nil {
		return nil, err
	}

	return deletedStation, nil
}

// FindStationByID retrieves a station by its ID
func (rs *StationService) GetAllStations() (*[]models.Station, error) {
	// Retrieve the station from the repository
	station, err := rs.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return station, nil
}

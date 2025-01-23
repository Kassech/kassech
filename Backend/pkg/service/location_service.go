package service

import (
	"fmt"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type LocationService struct {
	Repo *repository.LocationRepository
}

// CreateLocation creates a new location
func (ls *LocationService) CreateLocation(location *models.VehicleGPSLog) (*models.VehicleGPSLog, error) {
	createdLocation, err := ls.Repo.Create(location)
	if err != nil {
		return nil, err
	}
	return createdLocation, nil
}

func (ls *LocationService) DeleteLocationByID(locationID uint, forceDelete bool) (*models.VehicleGPSLog, error) {
	// Find the location to delete
	location, err := ls.Repo.FindByID(locationID)
	fmt.Println("location:", location)
	if err != nil {
		return nil, err
	}

	// Delete the location
	deletedLocation, err := ls.Repo.DeleteByID(locationID, forceDelete)
	if err != nil {
		return nil, err
	}

	return deletedLocation, nil
}

// FindLocationByID retrieves a location by its ID
func (ls *LocationService) FindLocationByID(locationID uint) (*models.VehicleGPSLog, error) {
	location, err := ls.Repo.FindByID(locationID)
	if err != nil {
		return nil, err
	}
	return location, nil
}

// GetAllLocations retrieves all locations
func (ls *LocationService) GetAllLocations(page, perPage int, search string, vehicleID uint, pathID uint) ([]models.VehicleGPSLog, int64, error) {
	// Implement pagination and search logic here
	locations, total, err := ls.Repo.GetAll(page, perPage, search, vehicleID, pathID)
	if err != nil {
		return nil, 0, err
	}
	return locations, total, nil
}

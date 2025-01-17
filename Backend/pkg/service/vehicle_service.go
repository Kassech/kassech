package service

import (
	"fmt"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type VehicleService struct {
	Repo *repository.VehicleRepository
}

func (vs *VehicleService) UpdateVehicle(vehicle *models.Vehicle, vehicleID uint) (*models.Vehicle, error) {
	// Perform the update operation
	updatedVehicle, err := vs.Repo.Update(vehicle, vehicleID)
	if err != nil {
		return nil, err
	}
	return updatedVehicle, nil
}

// CreateVehicle creates a new vehicle
func (vs *VehicleService) CreateVehicle(vehicle *models.Vehicle) (*models.Vehicle, error) {
	createdVehicle, err := vs.Repo.Create(vehicle)
	if err != nil {
		return nil, err
	}
	return createdVehicle, nil
}

func (vs *VehicleService) DeleteVehicleByID(vehicleID uint, forceDelete bool) (*models.Vehicle, error) {
	// Find the vehicle to delete
	vehicle, err := vs.Repo.FindByID(vehicleID)
	fmt.Println("vehicle:", vehicle)
	if err != nil {
		return nil, err
	}

	// Delete the vehicle
	deletedVehicle, err := vs.Repo.DeleteByID(vehicleID, forceDelete)
	if err != nil {
		return nil, err
	}

	return deletedVehicle, nil
}

// FindVehicleByID retrieves a vehicle by its ID
func (vs *VehicleService) FindVehicleByID(vehicleID uint) (*models.Vehicle, error) {
	vehicle, err := vs.Repo.FindByID(vehicleID)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

// GetAllVehicles retrieves all vehicles
func (vs *VehicleService) GetAllVehicles(page, perPage int, search, ownerID, typeID string) ([]models.Vehicle, int64, error) {
	// Implement pagination and search logic here
	vehicles, total, err := vs.Repo.GetAll(page, perPage, search, ownerID, typeID)
	if err != nil {
		return nil, 0, err
	}
	return vehicles, total, nil
}

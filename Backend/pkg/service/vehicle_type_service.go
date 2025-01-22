package service

import (
	"fmt"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type VehicleTypeService struct {
	Repo *repository.VehicleTypeRepository
}

func (vts *VehicleTypeService) UpdateVehicleType(vehicleType *models.VehicleType, vehicleTypeID uint) (*models.VehicleType, error) {
	// Perform the update operation
	updatedVehicleType, err := vts.Repo.Update(vehicleType, vehicleTypeID)
	if err != nil {
		return nil, err
	}
	return updatedVehicleType, nil
}

// CreateVehicleType creates a new vehicle type
func (vts *VehicleTypeService) CreateVehicleType(vehicleType *models.VehicleType) (*models.VehicleType, error) {
	createdVehicleType, err := vts.Repo.Create(vehicleType)
	if err != nil {
		return nil, err
	}
	return createdVehicleType, nil
}

func (vts *VehicleTypeService) DeleteVehicleTypeByID(vehicleTypeID uint, forceDelete bool) (*models.VehicleType, error) {
	// Find the vehicle type to delete
	vehicleType, err := vts.Repo.FindByID(vehicleTypeID,forceDelete)
	fmt.Println("vehicleType:", vehicleType)
	if err != nil {
		return nil, err
	}

	// Delete the vehicle type
	deletedVehicleType, err := vts.Repo.DeleteByID(vehicleTypeID, forceDelete)
	if err != nil {
		return nil, err
	}

	return deletedVehicleType, nil
}

// FindVehicleTypeByID retrieves a vehicle type by its ID
func (vts *VehicleTypeService) FindVehicleTypeByID(vehicleTypeID uint, forceDelete bool) (*models.VehicleType, error) {
	vehicleType, err := vts.Repo.FindByID(vehicleTypeID, forceDelete )
	if err != nil {
		return nil, err
	}
	return vehicleType, nil
}

// GetAllVehicleTypes retrieves all vehicle types
func (vts *VehicleTypeService) GetAllVehicleTypes() (*[]models.VehicleType, error) {
	vehicleTypes, err := vts.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return vehicleTypes, nil
}

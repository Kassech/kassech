package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type DriverDeligationService struct {
	Repo *repository.DriverDeligationRepository
}

func (s *DriverDeligationService) AssignDriverToVehicle(driverID, vehicleID uint) (*models.DriverVehicle, error) {
	assignment := &models.DriverVehicle{
		DriverID:  driverID,
		VehicleID: vehicleID,
		Status:    "active",
	}

	if err := s.Repo.Create(assignment); err != nil {
		return nil, err
	}
	return assignment, nil
}

func (s *DriverDeligationService) GetActiveDriver(vehicleID uint) (*models.DriverVehicle, error) {
	return s.Repo.GetActiveDriver(vehicleID)
}

package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type LocationService struct {
	locationRepo repository.LocationRepository
}

func NewLocationService(locationRepo repository.LocationRepository) *LocationService {
	return &LocationService{locationRepo: locationRepo}
}

func (s *LocationService) UpdateLocation(VehicleID uint, longitude, latitude float64, pathID uint) error {
	location := &models.VehicleGPSLog{
		VehicleID:    VehicleID,
		Latitude:  latitude,
		Longitude: longitude,
		PathID:     pathID,
	}

	_, err := s.locationRepo.Create(location)
	return err
}

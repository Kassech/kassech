package service

import (
	"kassech/backend/pkg/repository"
)

type PassengerService struct {
	passengerRepo repository.PassengerRepository
}

func NewPassengerService(passengerRepo repository.PassengerRepository) *PassengerService {
	return &PassengerService{passengerRepo: passengerRepo}
}

func (s *PassengerService) IncrementPassengerCountBy(pathID uint, amount int) error {
	return s.passengerRepo.IncrementPassengerCountBy(pathID, amount)
}

func (s *PassengerService) DecrementPassengerCountBy(pathID uint, amount int) error {
	return s.passengerRepo.DecrementPassengerCountBy(pathID, amount)
}

func (s *PassengerService) GetPassengerCount(pathID uint) (int64, error) {
	return s.passengerRepo.GetPassengerCount(pathID)
}

package service

import (
	"kassech/backend/pkg/repository"
)

type DestinationService struct {
	destinationRepo repository.DestinationRepository
}

func NewDestinationService(destinationRepo repository.DestinationRepository) *DestinationService {
	return &DestinationService{destinationRepo: destinationRepo}
}

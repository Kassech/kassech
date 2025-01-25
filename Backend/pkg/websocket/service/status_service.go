package service

import (
	"kassech/backend/pkg/repository"
)

type StatusService struct {
	userRepo repository.UserRepository
}

func NewStatusService(userRepo repository.UserRepository) *StatusService {
	return &StatusService{userRepo: userRepo}
}

func (s *StatusService) SetOnline(userID uint) error {
	return s.userRepo.UpdateUserStatus(userID, true)
}

func (s *StatusService) SetOffline(userID uint) error {
	return s.userRepo.UpdateUserStatus(userID, false)
}

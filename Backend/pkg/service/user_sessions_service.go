package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
	"time"
)

type UserSessionsService struct {
	Repo *repository.UserSessionsRepository
}

func (s *UserSessionsService) CreateSession(userID uint, token string, expiresAt time.Time) error {

	session := &models.UserSession{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return s.Repo.CreateSession(session)
}

func (s *UserSessionsService) GetOnlineUsers() ([]uint, error) {
	// Fetch unique online users by aggregating active sessions
	var onlineUsers []uint
	// Assuming repository method returns sessions grouped by user
	sessions, err := s.Repo.GetActiveSessionsByUserID(0) // Pass 0 for fetching all
	if err != nil {
		return nil, err
	}
	for _, session := range sessions {
		onlineUsers = append(onlineUsers, session.UserID)
	}
	return onlineUsers, nil
}

func (s *UserSessionsService) Logout(token string) error {
	return s.Repo.InvalidateSession(token)
}

func (s *UserSessionsService) DisableUser(userID string) error {
	// Invalidate all user sessions
	return s.Repo.InvalidateAllSessionsByUserID(userID)
}

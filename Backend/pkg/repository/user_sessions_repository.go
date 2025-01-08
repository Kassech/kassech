package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
	"time"
)

type UserSessionsRepository struct{}

func (r *UserSessionsRepository) CreateSession(session *models.UserSession) error {
	return database.DB.Create(session).Error
}

func (r *UserSessionsRepository) GetActiveSessionsByUserID(userID uint) ([]models.UserSession, error) {
	var sessions []models.UserSession
	err := database.DB.Where("user_id = ? AND is_active = TRUE AND expires_at > ?", userID, time.Now()).Find(&sessions).Error
	return sessions, err
}

func (r *UserSessionsRepository) InvalidateSession(token string) error {
	return database.DB.Model(&models.UserSession{}).Where("token = ?", token).Update("is_active", false).Error
}

func (r *UserSessionsRepository) InvalidateAllSessionsByUserID(userID string) error {
	return database.DB.Model(&models.UserSession{}).Where("user_id = ?", userID).Update("is_active", false).Error
}

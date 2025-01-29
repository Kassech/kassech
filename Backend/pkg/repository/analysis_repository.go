package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type AnalysisRepository struct{}

func (ar *AnalysisRepository) CountUsers() (int64, error) {
	var count int64
	err := database.DB.Model(&models.User{}).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountActiveUsers() (int64, error) {
	var count int64
	err := database.DB.Model(&models.User{}).Where("is_online = ?", true).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountDrivers() (int64, error) {
	var count int64
	err := database.DB.Model(&models.Driver{}).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountActiveDrivers() (int64, error) {
	var count int64
	err := database.DB.Model(&models.Driver{}).Where("status = ?", "active").Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountVehicles() (int64, error) {
	var count int64
	err := database.DB.Model(&models.Vehicle{}).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountActiveVehicles() (int64, error) {
	var count int64
	err := database.DB.Model(&models.Vehicle{}).Where("status = ?", "active").Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountRoutes() (int64, error) {
	var count int64
	err := database.DB.Model(&models.Route{}).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) CountTravelLogs() (int64, error) {
	var count int64
	err := database.DB.Model(&models.TravelLog{}).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) GetTotalStations() (int64, error) {
	var count int64
	err := database.DB.Model(&models.Station{}).Count(&count).Error
	return count, err
}

func (ar *AnalysisRepository) GetLoginLogs() ([]models.UserLoginLog, error) {
	var logs []models.UserLoginLog
	err := database.DB.Model(&models.UserLoginLog{}).Find(&logs).Error
	return logs, err
}

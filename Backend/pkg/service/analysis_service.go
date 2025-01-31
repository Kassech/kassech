package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type AnalysisService struct {
	Repo *repository.AnalysisRepository
}

func (as *AnalysisService) GetTotalUsers() (int64, error) {
	return as.Repo.CountUsers()
}

func (as *AnalysisService) GetActiveUsers() (int64, error) {
	return as.Repo.CountActiveUsers()
}

func (as *AnalysisService) GetTotalDrivers() (int64, error) {
	return as.Repo.CountDrivers()
}

func (as *AnalysisService) GetActiveDrivers() (int64, error) {
	return as.Repo.CountActiveDrivers()
}

func (as *AnalysisService) GetTotalVehicles() (int64, error) {
	return as.Repo.CountVehicles()
}

func (as *AnalysisService) GetActiveVehicles() (int64, error) {
	return as.Repo.CountActiveVehicles()
}

func (as *AnalysisService) GetTotalRoutes() (int64, error) {
	return as.Repo.CountRoutes()
}

func (as *AnalysisService) GetTotalTravelLogs() (int64, error) {
	return as.Repo.CountTravelLogs()
}

func (as *AnalysisService) GetTotalStations() (int64, error) {
	return as.Repo.GetTotalStations()
}

func (as *AnalysisService) GetLoginLogs() ([]models.UserLoginLog, error) {
	return as.Repo.GetLoginLogs()
}

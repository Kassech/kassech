package controller

import (
	responses "kassech/backend/pkg/response"
	"kassech/backend/pkg/service"

	"github.com/gin-gonic/gin"
)

type AnalysisController struct {
	Service *service.AnalysisService
}

// Define a constant response function for all endpoints

func (ac *AnalysisController) GetTotalUsers(c *gin.Context) {
	totalUsers, err := ac.Service.GetTotalUsers()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalUsers)
}

func (ac *AnalysisController) GetActiveUsers(c *gin.Context) {
	activeUsers, err := ac.Service.GetActiveUsers()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, activeUsers)
}

func (ac *AnalysisController) GetTotalDrivers(c *gin.Context) {
	totalDrivers, err := ac.Service.GetTotalDrivers()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalDrivers)
}

func (ac *AnalysisController) GetActiveDrivers(c *gin.Context) {
	activeDrivers, err := ac.Service.GetActiveDrivers()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, activeDrivers)
}

func (ac *AnalysisController) GetTotalVehicles(c *gin.Context) {
	totalVehicles, err := ac.Service.GetTotalVehicles()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalVehicles)
}

func (ac *AnalysisController) GetActiveVehicles(c *gin.Context) {
	activeVehicles, err := ac.Service.GetActiveVehicles()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, activeVehicles)
}

func (ac *AnalysisController) GetTotalRoutes(c *gin.Context) {
	totalRoutes, err := ac.Service.GetTotalRoutes()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalRoutes)
}

func (ac *AnalysisController) GetTotalTravelLogs(c *gin.Context) {
	totalTravelLogs, err := ac.Service.GetTotalTravelLogs()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalTravelLogs)
}
func (ac *AnalysisController) GetTotalStations(c *gin.Context) {
	totalStations, err := ac.Service.GetTotalStations()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalStations)
}

func (ac *AnalysisController) GetLoginLogs(c *gin.Context) {
	totalStations, err := ac.Service.GetLoginLogs()
	if err != nil {
		responses.BadRequest(c, err,
			responses.WithMessage("Invalid request parameters"),
		)
		return
	}
	responses.Success(c, totalStations)
}

package repository

import (
	"fmt"
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type RouteRepository struct{}

// Create a new route
func (rr *RouteRepository) Create(route *models.Route) (*models.Route, error) {
	fmt.Println("made it look", route)
	if err := database.DB.Create(route).Error; err != nil {
		fmt.Println("made it look", err)

		return nil, err
	}
	return route, nil
}

// Find a route by ID
func (rr *RouteRepository) FindByID(routeID uint) (*models.Route, error) {
	var route models.Route
	if err := database.DB.First(&route, routeID).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

// Get all routes
func (rr *RouteRepository) GetAll() (*[]models.Route, error) {
	var routes []models.Route
	if err := database.DB.Preload("StationA").Preload("StationB").Find(&routes).Error; err != nil {
		return nil, err
	}
	return &routes, nil
}

func (rr *RouteRepository) DeleteByLocation(locationA, locationB uint) error {
	return database.DB.Where("location_a = ? AND location_b = ?", locationA, locationB).Delete(&models.Route{}).Error
}

// Update an existing route by ID
func (rr *RouteRepository) Update(route *models.Route, routeID uint) (*models.Route, error) {
	// Check if the route exists
	existingRoute, err := rr.FindByID(routeID)
	if err != nil {
		return nil, err
	}

	// Use GORM's Updates method to update only the fields provided in the input
	if err := database.DB.Model(existingRoute).Updates(route).Error; err != nil {
		return nil, err
	}

	return existingRoute, nil
}

// Delete a route by ID
func (rr *RouteRepository) DeleteByID(routeID uint) (*models.Route, error) {
	var route models.Route
	// Find the route first to ensure it exists
	if err := database.DB.First(&route, routeID).Error; err != nil {
		return nil, err
	}

	// Delete the route after fetching it
	if err := database.DB.Delete(&route).Error; err != nil {
		return nil, err
	}

	return &route, nil
}

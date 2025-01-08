package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type RouteService struct {
	Repo *repository.RouteRepository
}

func (rs *RouteService) UpdateRoute(route *models.Route, routeID uint) (*models.Route, error) {
	// Perform the update operation
	updatedRoute, err := rs.Repo.Update(route, routeID) 
	if err != nil {
		return nil, err
	}
	return updatedRoute, nil
}

// CreateRoute creates a new route
func (rs *RouteService) CreateRoute(route *models.Route) (*models.Route, error) {
	// Create the primary route
	createdRoute, err := rs.Repo.Create(route)
	if err != nil {
		return nil, err
	}

	// Create the reverse route
	reverseRoute := &models.Route{
		LocationA: route.LocationB,
		LocationB: route.LocationA,
	}
	_, err = rs.Repo.Create(reverseRoute)
	if err != nil {
		return nil, err
	}

	return createdRoute, nil
}

func (rs *RouteService) DeleteRouteByID(routeID uint) (*models.Route, error) {
	// Find the route to delete
	route, err := rs.Repo.FindByID(routeID)
	if err != nil {
		return nil, err
	}

	// Delete the reverse route
	err = rs.Repo.DeleteByLocation(route.LocationB, route.LocationA)
	if err != nil {
		return nil, err
	}

	// Delete the primary route
	deletedRoute, err := rs.Repo.DeleteByID(routeID)
	if err != nil {
		return nil, err
	}

	return deletedRoute, nil
}

// FindRouteByID retrieves a route by its ID
func (rs *RouteService) FindRouteByID(routeID uint) (*models.Route, error) {
	route, err := rs.Repo.FindByID(routeID)
	if err != nil {
		return nil, err
	}
	return route, nil
}

// GetAllRoutes retrieves all routes
func (rs *RouteService) GetAllRoutes() (*[]models.Route, error) {
	routes, err := rs.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return routes, nil
}

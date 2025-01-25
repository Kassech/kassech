package service

import (
	"errors"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type QueueManagerRouteService struct {
	Repo *repository.QueueManagerRouteRepository
}

type CreateQueueManagerRouteRequest struct {
	UserID    uint   `json:"userId" binding:"required"`
	StationID uint   `json:"stationId" binding:"required"`
	PathIDs   []uint `json:"pathIds" binding:"required,min=1"`
}

func (s *QueueManagerRouteService) CreateRoute(req CreateQueueManagerRouteRequest) (*models.QueueManagerRoute, error) {
	if len(req.PathIDs) == 0 {
		return nil, errors.New("at least one path must be selected")
	}

	route := &models.QueueManagerRoute{
		UserID:    req.UserID,
		StationID: req.StationID,
	}

	if err := s.Repo.Create(route, req.PathIDs); err != nil {
		return nil, err
	}

	return route, nil
}
func (s *QueueManagerRouteService) GetAllRoutes(page, perPage int, search string) ([]models.QueueManagerRoute, int64, error) {
	return s.Repo.GetAll(page, perPage, search)
}

func (s *QueueManagerRouteService) DeleteRoute(id uint) error {
	return s.Repo.Delete(id)
}

func (s *QueueManagerRouteService) GetAllForUser(userID uint, page, perPage int, search string) ([]models.QueueManagerRoute, int64, error) {
	return s.Repo.GetAllForUser(userID, page, perPage, search)
}


func (s *QueueManagerRouteService) GetAllRoutesForUser(userID uint, page, perPage int, search string) ([]models.QueueManagerRoute, int64, error) {
	return s.Repo.GetAllForUser(userID, page, perPage, search)
}

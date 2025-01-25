package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type PathService struct {
	Repo *repository.PathRepository
}

func (ps *PathService) CreatePath(path *models.Path) (*models.Path, error) {
	createdPath, err := ps.Repo.Create(path)
	if err != nil {
		return nil, err
	}

	return createdPath, nil
}

func (ps *PathService) DeletePathByID(pathID uint) (*models.Path, error) {
	_, err := ps.Repo.FindByID(pathID)
	if err != nil {
		return nil, err
	}

	deletedPath, err := ps.Repo.DeleteByID(pathID)
	if err != nil {
		return nil, err
	}

	return deletedPath, nil
}

func (ps *PathService) FindPathByID(pathID uint) (*models.Path, error) {
	path, err := ps.Repo.FindByID(pathID)
	if err != nil {
		return nil, err
	}
	return path, nil
}

func (ps *PathService) GetAllPaths(page, perPage int, search string) (*[]models.Path, int64, error) {
	paths, total, err := ps.Repo.GetAll(page, perPage, search)
	if err != nil {
		return nil, 0, err
	}
	return &paths, total, nil
}

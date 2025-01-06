package service

import (
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
)

type UserRoleService struct {
	Repo *repository.UserRoleRepository
}

func (s *UserRoleService) Create(userRole *models.UserRole) error {
	if err := userRole.Validate(); err != nil {
		return err
	}
	return s.Repo.Create(userRole)
}

func (s *UserRoleService) GetByID(id string) (*models.UserRole, error) {
	return s.Repo.FindByID(id)
}

func (s *UserRoleService) Update(id string, userRole *models.UserRole) error {
	existing, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	existing.UserID = userRole.UserID
	existing.RoleID = userRole.RoleID

	return s.Repo.Update(existing)
}

func (s *UserRoleService) Delete(id string) error {
	return s.Repo.Delete(id)
}

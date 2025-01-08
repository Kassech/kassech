package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type UserRoleRepository struct {
}

func (repo *UserRoleRepository) Create(userRole *models.UserRole) error {
	return database.DB.Create(userRole).Error
}

func (repo *UserRoleRepository) FindByID(id string) (*models.UserRole, error) {
	var userRole models.UserRole
	if err := database.DB.First(&userRole, id).Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (repo *UserRoleRepository) Update(userRole *models.UserRole) error {
	return database.DB.Save(userRole).Error
}

func (repo *UserRoleRepository) Delete(id string) error {
	return database.DB.Delete(&models.UserRole{}, id).Error
}

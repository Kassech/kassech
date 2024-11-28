package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type UserRepository struct{}

func (ur *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (ur *UserRepository) FindByEmailOrPhone(email, phone string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ? OR phone_number = ?", email, phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

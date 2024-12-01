package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type UserRepository struct{}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	// Attempt to create the user in the database
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	// Return the created user along with nil error if the creation was successful
	return user, nil
}

func (ur *UserRepository) FindByEmailOrPhone(email string, phone string) (*models.User, error) {
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

func (ur *UserRepository) LogLoginEvent(loginLog *models.UserLoginLog) error {
	// Insert the new login event into the database
	if err := database.DB.Create(loginLog).Error; err != nil {
		return err
	}
	return nil
}

// Update updates the user's last login time in the database
func (ur *UserRepository) Update(user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

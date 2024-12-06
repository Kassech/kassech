package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type UserRepository struct{}

// Create inserts a new user into the database
func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmailOrPhone searches for a user by either email or phone number
func (ur *UserRepository) FindByEmailOrPhone(email string, phone string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ? OR phone_number = ?", email, phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID fetches a user by their unique ID
func (ur *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// LogLoginEvent inserts a new login event for the user
func (ur *UserRepository) LogLoginEvent(loginLog *models.UserLoginLog) error {
	if err := database.DB.Create(loginLog).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing user in the database
func (ur *UserRepository) Update(user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// Delete removes a user from the database
func (ur *UserRepository) Delete(user *models.User) error {
	if err := database.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

// ListUsers fetches users with pagination and optional search filter
func (ur *UserRepository) ListUsers(page, limit int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Build the query with optional search filter
	query := database.DB.Model(&models.User{})
	if search != "" {
		// Search by name, email, or phone number
		query = query.Where("email LIKE ? OR phone_number LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Get the total count of users matching the search
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Fetch the users with pagination (offset and limit)
	err = query.Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

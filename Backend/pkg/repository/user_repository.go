package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
)

type UserRepository struct{}

// Create inserts a new user and their role into the database
func (ur *UserRepository) Create(user *models.User, role uint) (*models.User, error) {
	tx := database.DB.Begin()

	// Insert the user
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create and associate the user role
	userRole := &models.UserRole{
		UserID: user.ID,
		RoleID: role,
	}
	if err := tx.Create(userRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	tx.Commit()

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

// ListUsers fetches users with pagination, optional search filter, and active/deleted filter
func (ur *UserRepository) ListUsers(page, limit int, search string, typ string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Build the base query
	query := database.DB.Table("users").
		Select("users.id, users.first_name, users.last_name, users.email, users.phone_number, COALESCE(array_agg(roles.role_name), '{}') AS roles").
		Joins("LEFT JOIN user_roles ON user_roles.user_id = users.id").
		Joins("LEFT JOIN roles ON roles.id = user_roles.role_id")

	// Apply the search filter
	if search != "" {
		query = query.Where("email LIKE ? OR phone_number LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Filter by the 'typ' parameter (active or deleted)
	switch typ {
	case "active":
		query = query.Where("users.deleted_at IS NULL")
	case "deleted":
		query = query.Unscoped().Where("users.deleted_at IS NOT NULL")
	default:
		query = query.Where("users.deleted_at IS NULL")
	}

	// Get the total number of users matching the filters
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Retrieve the users with roles and pagination using raw SQL query
	err = query.Offset((page - 1) * limit).Limit(limit).Group("users.id").Scan(&users).Error
	if err != nil {
		return nil, 0, err
	}

	// Convert roles from a comma-separated string to a slice of strings
	for i := range users {
		roles := users[i].Roles
		users[i].Roles = roles // Split the comma-separated string into a slice of strings
	}

	return users, total, nil
}

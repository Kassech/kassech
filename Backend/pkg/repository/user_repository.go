package repository

import (
	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"

	"gorm.io/gorm"
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

// CreateDriver inserts a new driver
func (ur *UserRepository) CreateDriver(driver *models.Driver) (*models.Driver, error) {
	err := database.DB.Create(driver).Error
	if err != nil {
		return nil, err
	}
	return driver, nil
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
func (ur *UserRepository) Delete(user *models.User, isForce bool) error {
	query := database.DB
	if isForce {
		query = query.Unscoped()
	}
	return query.Delete(user).Error
}

// ListUsers fetches users with pagination, optional search filter, and active/deleted filter
func (ur *UserRepository) ListUsers(page, limit int, search, typ, role string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Build the base query
	query := database.DB.
		Model(&models.User{}).
		Select("users.id, users.first_name, users.last_name, users.email, users.phone_number, array_agg(DISTINCT roles.role_name) AS roles").
		Joins("LEFT JOIN user_roles ON user_roles.user_id = users.id").
		Joins("LEFT JOIN roles ON roles.id = user_roles.role_id").
		Group("users.id")

	// Apply the search filter
	if search != "" {
		query = query.Where("email ILIKE ? OR phone_number ILIKE ?", "%"+search+"%", "%"+search+"%")
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

	// Filter by the 'role' parameter
	if role != "" {
		query = query.Where("roles.id = ?", role)
	}

	// Get the total number of users matching the filters
	err := query.Count(&total).Error
	if err != nil {
		return []models.User{}, 0, err
	}

	// Retrieve the users with roles and pagination
	err = query.Offset((page - 1) * limit).Limit(limit).Scan(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
func (ur *UserRepository) SaveNotificationToken(userID uint, token string, deviceID string) error {
	var existingToken models.NotificationToken

	// Check if a token with the same device ID already exists
	err := database.DB.Where("user_id = ? AND device_id = ?", userID, deviceID).First(&existingToken).Error
	if err == nil {
		// Update the existing token
		existingToken.Token = token
		existingToken.Status = "active"
		if err := database.DB.Save(&existingToken).Error; err != nil {
			return err
		}
	} else if err == gorm.ErrRecordNotFound {
		// Insert a new token
		notificationToken := &models.NotificationToken{
			UserID:   userID,
			Token:    token,
			Status:   "active",
			DeviceID: deviceID,
		}
		if err := database.DB.Create(notificationToken).Error; err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

// GetPermissionsAndRolesByUserID gets a list of permissions and roles based on a user ID
func (ur *UserRepository) GetPermissionsAndRolesByUserID(userID uint) ([]string, []string, error) {
	var permissions []struct {
		PermissionName string
	}
	var roles []struct {
		RoleName string
	}

	err := database.DB.Model(&models.UserRole{}).
		Select("DISTINCT p.permission_name").
		Joins("JOIN user_roles AS ur2 ON ur2.user_id = ?", userID).
		Joins("JOIN role_permissions AS rp ON rp.role_id = ur2.role_id").
		Joins("JOIN permissions AS p ON p.id = rp.permission_id").
		Find(&permissions).Error
	if err != nil {
		return nil, nil, err
	}

	err = database.DB.Model(&models.UserRole{}).
		Select("DISTINCT r.role_name").
		Joins("JOIN user_roles AS ur2 ON ur2.user_id = ?", userID).
		Joins("JOIN roles AS r ON r.id = ur2.role_id").
		Find(&roles).Error
	if err != nil {
		return nil, nil, err
	}

	permissionList := make([]string, len(permissions))
	roleList := make([]string, len(roles))
	for i, p := range permissions {
		permissionList[i] = p.PermissionName
	}
	for i, r := range roles {
		roleList[i] = r.RoleName
	}

	return permissionList, roleList, nil
}

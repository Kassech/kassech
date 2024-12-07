package service

import (
	"errors"
	"log"
	"net/http"
	"time"

	"kassech/backend/pkg/database"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

// Register a new user
func (us *UserService) Register(user *models.User) (*models.User, string, string, error) {
	if err := user.Validate(); err != nil {
		return nil, "", "", err
	}

	existingUser, _ := us.Repo.FindByEmailOrPhone(user.Email, user.PhoneNumber)
	if existingUser != nil {
		// User already exists
		return nil, "", "", errors.New("user with that email or phone number already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	user, err = us.Repo.Create(user)
	if err != nil {
		return nil, "", "", err
	}

	// Generate the JWT tokens
	accessToken, refreshToken, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to generate token")
	}

	// Log the registration login event
	us.LogLoginEvent(user, nil) // No IP or UserAgent needed during registration

	return user, accessToken, refreshToken, nil
}


// Create  a new user with Specific Role and data 
func (us *UserService) CreateUser(user *models.User) (*models.User, string, string, error) {
	if err := user.Validate(); err != nil {
		return nil, "", "", err
	}

	existingUser, _ := us.Repo.FindByEmailOrPhone(user.Email, user.PhoneNumber)
	if existingUser != nil {
		// User already exists
		return nil, "", "", errors.New("user with that email or phone number already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	user, err = us.Repo.Create(user)
	if err != nil {
		return nil, "", "", err
	}

	// Generate the JWT tokens
	accessToken, refreshToken, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to generate token")
	}

	// Log the registration login event
	us.LogLoginEvent(user, nil) // No IP or UserAgent needed during registration

	return user, accessToken, refreshToken, nil
}
// Login handles the user login
func (us *UserService) Login(emailOrPhone, password string, r *http.Request) (*models.User, string, string, error) {
	user, err := us.Repo.FindByEmailOrPhone(emailOrPhone, emailOrPhone)
	if err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Compare password hash with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Generate the JWT tokens
	accessToken, refreshToken, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", "", errors.New("failed to generate token")
	}

	// Log the login event
	us.LogLoginEvent(user, r)

	return user, accessToken, refreshToken, nil
}

// LogLoginEvent logs the login event for a user
func (us *UserService) LogLoginEvent(user *models.User, r *http.Request) {
	// Default values for IP address and UserAgent
	var ipAddress string
	var userAgent string

	// Check if r is not nil
	if r != nil {
		// Try to capture the client IP address from the X-Forwarded-For header
		ipAddress = r.Header.Get("X-Forwarded-For")

		// If X-Forwarded-For is not set, fall back to RemoteAddr
		if ipAddress == "" {
			ipAddress = r.RemoteAddr
		}

		// If RemoteAddr is also empty, set ipAddress to an empty string
		if ipAddress == "" {
			ipAddress = "" // Or use nil if using pointer
		}

		// Capture the user agent from the request
		userAgent = r.UserAgent()
	}

	// Create the UserLoginLog entry
	loginLog := models.UserLoginLog{
		UserID:    user.ID,
		IP:        ipAddress,
		UserAgent: userAgent,
		LoginTime: time.Now(),
	}

	// Store the login log in the database
	err := us.Repo.LogLoginEvent(&loginLog)
	if err != nil {
		log.Printf("Error logging login event: %v\n", err)
		// Continue without returning the error
	}

	// Log the login event (optional for debugging)
	log.Println(loginLog)

	// Update the user's last login time (for user management purposes)
	user.LastLoginDate = time.Now()
	err = us.Repo.Update(user)
	if err != nil {
		log.Printf("Error updating user last login date: %v\n", err)
		// Continue without returning the error
	}
}

// ListUsers with Pagination and Search
func (us *UserService) ListUsers(page, limit int, search string, typ string) ([]models.User, int64, error) {
	// Define pagination parameters
	offset := (page - 1) * limit
	var users []models.User
	var total int64

	// Build the query with optional search filter
	query := database.DB.Model(&models.User{})

	// Apply the search filter (search by email or phone number)
	if search != "" {
		query = query.Where("email LIKE ? OR phone_number LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Filter by the 'typ' parameter (active or deleted)
	switch typ {
	case "active":
		// For active users, we want to check that 'deleted_at' is NULL
		query = query.Where("deleted_at IS NULL")

	case "deleted":
		// For deleted users, we want to include those with a non-NULL 'deleted_at'
		query = query.Unscoped().Where("deleted_at IS NOT NULL")

	default:
		// Default behavior: If 'typ' is not provided or an unknown value, treat it as 'active'
		query = query.Where("deleted_at IS NULL")
	}

	// Get the total number of users based on the query with filters
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Retrieve users with pagination
	err = query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateUser updates a user by ID
func (us *UserService) UpdateUser(userId string, user *models.User) (*models.User, error) {
	existingUser, err := us.Repo.FindByID(utils.StringToUint(userId))
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update user fields
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.PhoneNumber = user.PhoneNumber

	// Save the updated user
	err = us.Repo.Update(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(userId string) error {
	existingUser, err := us.Repo.FindByID(utils.StringToUint(userId))
	if err != nil {
		return errors.New("user not found")
	}

	// Delete the user
	err = us.Repo.Delete(existingUser)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"kassech/backend/pkg/domain"
	"kassech/backend/pkg/mapper"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (us *UserService) CreateUser(user *models.User, role uint) (*models.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	existingUser, _ := us.Repo.FindByEmailOrPhone(user.Email, user.PhoneNumber)
	if existingUser != nil {
		// User already exists
		return nil, errors.New("user with that email or phone number already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Println(user.Password)
	fmt.Println(string(hashedPassword))
	if err != nil {
		log.Println("Password hashing error:", err)       // Log the specific bcrypt error
		return nil, errors.New("failed to hash password") // Return a generic error to the user
	}

	user.Password = string(hashedPassword)

	// Create the user in the database
	user, err = us.Repo.Create(user, role)
	if err != nil {
		return nil, err
	}
	// Log the registration login event
	us.LogLoginEvent(user, nil) // No IP or UserAgent needed during registration

	return user, nil
}

// GenerateAuthentication generates an access token and a refresh token for a user
func (us *UserService) GenerateAuthentication(user *models.User) (*domain.User, string, string, error) {
	// Make sure the user exists
	existingUser, err := us.Repo.FindByEmailOrPhone(user.Email, user.PhoneNumber)
	if err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Generate the JWT tokens
	domainUser := mapper.ToDomainUser(existingUser)
	userPermissions, userRole, err := us.Repo.GetPermissionsAndRolesByUserID(user.ID)
	fmt.Println("userRole:", userRole)
	fmt.Println("userPermissions:", userPermissions)
	if err != nil {
		domainUser.Permissions = []string{}
	} else {
		domainUser.Permissions = userPermissions
	}

	accessToken, refreshToken, err := GenerateToken(user.ID, userRole)
	if err != nil {
		return nil, "", "", errors.New("failed to generate token")
	}

	return domainUser, accessToken, refreshToken, nil
}

// SaveNotificationToken saves the notification token for a user
func (us *UserService) SaveNotificationToken(userID uint, token string, device_id string) error {
	// Call the repository method to save the notification token
	err := us.Repo.SaveNotificationToken(userID, token, device_id)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) VerifyUser(userID uint, state bool) (*models.User, error) {
	// Call the repository method to save the notification token
	user, err := us.Repo.VerifyUser(userID, state)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login handles the user login
func (us *UserService) Login(emailOrPhone, password string, r *http.Request) (*domain.User, string, string, error) {
	user, err := us.Repo.FindByEmailOrPhone(emailOrPhone, emailOrPhone)
	if err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Compare password hash with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Convert the model to domain
	domainUser := mapper.ToDomainUser(user)
	userPermissions, userRole, err := us.Repo.GetPermissionsAndRolesByUserID(user.ID)

	if err != nil {
		domainUser.Permissions = []string{}
		domainUser.Roles = []string{}
	} else {
		domainUser.Permissions = userPermissions
		domainUser.Roles = userRole
	}

	// Generate the JWT tokens
	accessToken, refreshToken, err := GenerateToken(user.ID, userRole)

	if err != nil {
		return nil, "", "", errors.New("failed to generate token")
	}

	// Log the login event
	us.LogLoginEvent(user, r)

	return domainUser, accessToken, refreshToken, nil
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

func (us *UserService) ListUsers(page, limit int, search string, typ string, role string) ([]models.User, int64, error) {
	// Call the repository method
	return us.Repo.ListUsers(page, limit, search, typ, role)
}

// UpdateUser updates a user by ID
// UpdateUser updates a user's information and related driver documents if applicable
func (us *UserService) UpdateUser(userId uint, user *models.User) (*models.User, error) {
	existingUser, err := us.Repo.FindByID(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields if provided
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.PhoneNumber != "" {
		existingUser.PhoneNumber = user.PhoneNumber
	}
	if user.ProfilePicture != nil {
		existingUser.ProfilePicture = user.ProfilePicture
	}

	// Handle password update
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		existingUser.Password = string(hashedPassword)
	}

	// Save the updated user
	if err := us.Repo.Update(existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

// GetDriverByUserID retrieves a driver by their associated user ID
func (us *UserService) GetDriverByUserID(userID uint) (*models.Driver, error) {
	return us.Repo.GetDriverByUserID(userID)
}

// UpdateDriver updates a driver's document paths
func (us *UserService) UpdateDriver(driver *models.Driver) (*models.Driver, error) {
	return us.Repo.UpdateDriver(driver)
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(userId uint, isForce ...bool) error {
	force := len(isForce) > 0 && isForce[0]

	log.Printf("Deleting user with ID: %v (force: %v)\n", userId, force)

	existingUser, err := us.Repo.FindByID(userId)
	if err != nil {
		log.Printf("Error finding user with ID: %v\n", userId)
		return errors.New("user not found")
	}

	log.Printf("Found user with ID: %v\n", userId)

	// Delete the user
	err = us.Repo.Delete(existingUser, force)
	if err != nil {
		log.Printf("Error deleting user with ID: %v\n", userId)
		return err
	}

	log.Printf("Deleted user with ID: %v\n", userId)

	return nil
}

func (us *UserService) CreateDriver(driver *models.Driver) (*models.Driver, error) {
	return us.Repo.CreateDriver(driver)
}

// GetUserById fetches a user by their unique ID
func (us *UserService) GetUserById(userId uint) (*models.User, error) {
	user, err := us.Repo.FindByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

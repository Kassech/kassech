package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"
	"kassech/backend/pkg/utils"

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
	if err != nil {
		return nil, errors.New("failed to hash password")
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

// Register a new user
func (us *UserService) Register(user *models.User, role uint) (*models.User, string, string, error) {
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
	user, err = us.Repo.Create(user, role)

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

// SaveNotificationToken saves the notification token for a user
func (us *UserService) SaveNotificationToken(userID uint, token string, device_id string) error {
	// Call the repository method to save the notification token
	err := us.Repo.SaveNotificationToken(userID, token, device_id)
	if err != nil {
		return err
	}

	return nil
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
	fmt.Println("accessToken, refreshToken, err:", accessToken, refreshToken, err)

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

func (us *UserService) ListUsers(page, limit int, search string, typ string) ([]models.User, int64, error) {
	// Call the repository method
	return us.Repo.ListUsers(page, limit, search, typ)
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

package service

import (
	"errors"
	"log"
	"net/http"
	"time"

	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"

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
func (us *UserService) LogLoginEvent(user *models.User, r *http.Request) error {
	// Capture the IP address from the request
	ipAddress := r.RemoteAddr
	// Optionally, you could parse the IP address from a proxy (if you are behind a proxy, e.g., in production)
	// ipAddress := r.Header.Get("X-Forwarded-For") // Uncomment if behind a proxy

	// Capture the user agent from the request
	userAgent := r.UserAgent()

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
		return err
	}
	log.Println(loginLog)
	// Update the user's last login time (for user management purposes)
	user.LastLoginDate = time.Now()
	err = us.Repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

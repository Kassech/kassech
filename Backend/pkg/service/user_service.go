package service

import (
	"errors"
	models "kassech/backend/pkg/model"
	"kassech/backend/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (us *UserService) Register(user *models.User) (*models.User, string, error) {
	if err := user.Validate(); err != nil {
		return nil, "", err
	}
	existingUser, _ := us.Repo.FindByEmailOrPhone(*user.Email, user.PhoneNumber)
	if existingUser != nil {
		// User already exists, handle edge case more gracefully
		return nil, "", errors.New("user with that email or phone number already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	user, err = us.Repo.Create(user)
	if err != nil {
		return nil, "", err
	}

	// Generate the JWT token
	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func (us *UserService) Login(emailOrPhone, password string) (*models.User, string, error) {
	user, err := us.Repo.FindByEmailOrPhone(emailOrPhone, emailOrPhone)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

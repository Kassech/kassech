package repository

import (
	"context"
	"errors"
	"kassech/backend/pkg/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// GetUserByEmail fetches a user by their email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	result := r.db.WithContext(ctx).Create(user)
	return result.Error
}

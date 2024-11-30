package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uint `validate:"required"`
	RoleID uint `validate:"required"`
}

// Validate method for UserRole
func (ur *UserRole) Validate() error {
	validate := validator.New()
	return validate.Struct(ur)
}

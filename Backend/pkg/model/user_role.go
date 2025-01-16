package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" validate:"required"`
	RoleID uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoleID;references:ID" validate:"required"`
}

// Validate method for UserRole
func (ur *UserRole) Validate() error {
	validate := validator.New()
	return validate.Struct(ur)
}

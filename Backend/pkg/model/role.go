package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName    string       `gorm:"not null" validate:"required"`
	Description string       `validate:"omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []User       `gorm:"many2many:user_roles;"`
}

// Validate method for Role
func (r *Role) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

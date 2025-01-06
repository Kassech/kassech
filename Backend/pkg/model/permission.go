package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	PermissionName string `gorm:"not null" validate:"required"`
	Description    string `validate:"omitempty"`
	Roles          []Role `gorm:"many2many:role_permissions;"`
}

// Validate method for Permission
func (p *Permission) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

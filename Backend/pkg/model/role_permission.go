package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleID       uint `validate:"required"`
	PermissionID uint `validate:"required"`
}

// Validate method for RolePermission
func (rp *RolePermission) Validate() error {
	validate := validator.New()
	return validate.Struct(rp)
}

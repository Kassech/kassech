package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string `gorm:"not null"`
	Description string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []User       `gorm:"many2many:user_roles;"`
}

package models

import "gorm.io/gorm"

type Permission struct {
    gorm.Model
    PermissionName string `gorm:"not null"`
    Description    string
}

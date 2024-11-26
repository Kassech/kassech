package models

import "gorm.io/gorm"

type NotificationToken struct {
	gorm.Model
	UserID uint   `gorm:"not null"` // Foreign key to User
	Token  string `gorm:"size:255;not null;unique"`
	Status string `gorm:"size:20;not null;default:'active'"`
}

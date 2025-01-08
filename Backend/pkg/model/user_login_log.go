package models

import (
	"time"

	"gorm.io/gorm"
)

type UserLoginLog struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	IP        string    `gorm:"not null"`
	UserAgent string    `gorm:"not null"`
	LoginTime time.Time `gorm:"not null"`
}

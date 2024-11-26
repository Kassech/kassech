package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	PhoneNumber    string `gorm:"size:20;not null"`
	IsOnline       bool   `gorm:"default:false"`
	Password       string `gorm:"not null"`
	ProfilePicture string
	IsVerified     bool `gorm:"default:false"`
}

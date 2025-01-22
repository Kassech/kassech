package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName         string `gorm:"not null" validate:"required"`
	LastName          string `gorm:"not null" validate:"required"`
	Email             string `gorm:"unique;default:null" validate:"omitempty,email"`
	PhoneNumber       string `gorm:"size:20;not null;unique" validate:"required,e164"`
	IsOnline          bool   `gorm:"default:false"`
	Password          string `gorm:"not null" validate:"required,min=6" json:"-"`
	PreviousPassword1 string `gorm:"size:255" json:"-"`
	PreviousPassword2 string `gorm:"size:255" json:"-"`
	ProfilePicture    *string
	IsVerified        bool      `gorm:"default:false"`
	LastLoginDate     time.Time `gorm:"default:null"`
	SocketID          *string   `gorm:"size:255;default:null"`
	Roles             string    `json:"roles"`
}

// Validate will use the validator package to validate the User struct
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

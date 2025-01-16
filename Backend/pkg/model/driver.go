// models/driver.go

package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	UserID             uint   `gorm:"not null"`
	User               User   `gorm:"foreignKey:UserID"` // Link to the User model
	DriverLicense      string `gorm:"not null"`
	Status             string
	DrivingLicensePath string
	NationalIdPath     string
	InsuranceDocPath   string
	OtherFilePath      string
}

func (d Driver) Validate() error {
	v := validator.New()
	return v.Struct(d)
}

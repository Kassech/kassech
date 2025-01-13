// models/driver.go

package models

import "gorm.io/gorm"

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

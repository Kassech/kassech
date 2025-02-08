// models/driver.go
package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	UserID             uint   `gorm:"not null;uniqueIndex"`
	User               User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	DriverLicense      string `gorm:"not null;size:50"`
	Status             string `gorm:"size:20;default:'pending'"`
	DrivingLicensePath string `gorm:"not null"`
	NationalIdPath     string `gorm:"not null"`
	InsuranceDocPath   string
	OtherFilePath      string
	Vehicle            *Vehicle `gorm:"foreignKey:DriverID;constraint:OnDelete:SET NULL"` // One-to-one back reference

}

func (d Driver) Validate() error {
	v := validator.New()
	return v.Struct(d)
}

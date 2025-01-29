package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	TypeID        uint
	Type          VehicleType `gorm:"foreignKey:TypeID;references:ID"`
	LicenseNumber string      `gorm:"size:20;not null"`
	VIN           string      `gorm:"size:20;unique;not null"`
	Make          string      `gorm:"not null"`
	Year          uint        `gorm:"not null"`
	Color         string
	CarPicture    string `gorm:"type:text"`
	Bollo         string `gorm:"type:text"`
	Insurance     string `gorm:"type:text"`
	Libre         string `gorm:"type:text"`
	OwnerID       uint   `gorm:"not null"`
	Status        string `gorm:"size:20;not null;default:'active'"`
	Owner         User   `gorm:"foreignKey:OwnerID;references:ID"`
}

func (v *Vehicle) SetCarPicture(path string) {
	v.CarPicture = path
}

func (v *Vehicle) SetBollo(path string) {
	v.Bollo = path
}

func (v *Vehicle) SetInsurance(path string) {
	v.Insurance = path
}

func (v *Vehicle) SetLibre(path string) {
	v.Libre = path
}

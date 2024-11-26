package models

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	Type  string `gorm:"not null"`
	Image string `gorm:"not null"`
}

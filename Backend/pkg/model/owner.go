package models

import "gorm.io/gorm"

type Owner struct {
    gorm.Model
    UserID uint
    CarID  uint
    Status string
}

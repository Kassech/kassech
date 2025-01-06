package models

import "gorm.io/gorm"

type NotificationLog struct {
    gorm.Model
    SenderID  uint
    ReceiverID uint
    Title     string `gorm:"not null"`
    Body      string `gorm:"not null"`
    ImageURL  string
    Status    string `gorm:"not null"`
}

package models

import (
	"gorm.io/gorm"
)

// User models
type User struct {
	gorm.Model
	Username         string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email            string  `gorm:"type:varchar(100);uniqueIndex;"` // Разрешаем пустые значения, но сохраняе
	Password         string  `gorm:"type:varchar(255);default:null"`
	PhoneNumber      string  `gorm:"unique;not null"`
	VerificationCode *string `gorm:"size:6"`
	ChatID           *string `gorm:"size:100"`
}

package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID            uint     `gorm:"primaryKey"`
	Address       string   `gorm:"type:text"`
	PaymentMethod *string  `gorm:"size:50"`
	UserID        uint     `gorm:"column:user_id;uniqueIndex"`
	User          User     `gorm:"foreignKey:UserID"`
	CompanyID     *uint    `gorm:"column:company_id"`
	Company       *Company `gorm:"foreignKey:CompanyID"`
}

package models

import (
	"gorm.io/gorm"
)

// Driver represents the driver table
type Driver struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey"`
	Avatar            string `gorm:"size:100"`
	PartnerType       string `gorm:"size:20"`
	Status            string `gorm:"size:20"`
	Rating            float32
	CarID             *uint    `gorm:"column:car_id"`
	PartnerID         *uint    `gorm:"column:partner_id"`
	UserID            uint     `gorm:"column:user_id;uniqueIndex"`
	CourierVariantID  *uint    `gorm:"column:courier_variant_id"`
	Services          string   `gorm:"size:50"`
	Car               *Car     `gorm:"foreignKey:CarID"`
	Partner           *Company `gorm:"foreignKey:PartnerID"`
	User              User     `gorm:"foreignKey:UserID"`
	DocumentsProvided bool
}

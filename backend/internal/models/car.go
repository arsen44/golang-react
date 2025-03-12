package models

import "gorm.io/gorm"

// DeliveryCar represents the delivery_car table
type Car struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	CarBrand      string `gorm:"size:50"`
	NumberPlate   string `gorm:"size:20"`
	SeatNumber    string `gorm:"size:20"`
	PhotoDocument string `gorm:"size:100"`
	UserID        *uint  `gorm:"column:user_id"`
	User          *User  `gorm:"foreignKey:UserID"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

// Role represents role company.
type Role string

// Possible role company.
const (
	CompanyPark Role = "Company/Park"
	Сlinet      Role = "Сlinet"
)

type Company struct {
	gorm.Model
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"size:255"`
	ContactEmail *string `gorm:"size:254"`
	Phone        *string `gorm:"size:20"`
	Address      *string `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role         *Role    `gorm:"size:20"`
	Commission   *float64 `gorm:"type:decimal(5,2)"`
}

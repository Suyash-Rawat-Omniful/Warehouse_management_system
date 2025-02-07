package models

import (
	"time"

	"gorm.io/gorm"
)

// SKU represents a product SKU in the system
type SKU struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	ProductID uint   `gorm:"not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	Price     int    `gorm:"not null"`
	Fragile   bool   `gorm:"not null"`
	ImageURL  string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

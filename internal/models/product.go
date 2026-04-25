package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  uint   `gorm:"index;not null"`
	Name      string `gorm:"not null"`
	SKU       string `gorm:"uniqueIndex"`
	Price     float64
	Cost      float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

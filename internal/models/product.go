package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  uint   `gorm:"index:idx_sku_tenant,unique;not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	SKU       string `gorm:"type:varchar(100);index:idx_sku_tenant,unique"`
	Price     float64
	Cost      float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

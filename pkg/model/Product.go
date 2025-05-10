package model

import "time"

type Product struct {
	ID            uint   `gorm:"primaryKey"`
	TenantID      uint   `gorm:"index;not null"`                                                   // 🔹 FK ไป Tenant.ID
	Tenant        Tenant `gorm:"foreignKey=TenantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // ✅ เพิ่มตรงนี้
	ProductCode   string `gorm:"uniqueIndex"`
	SKU           string `gorm:"index"`
	Name          string
	UOM           string
	PriceStandard float64
	VatCategoryID uint
	VatCategory   VatCategory
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

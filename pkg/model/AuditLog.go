package model

import (
	"time"

	"gorm.io/datatypes"
)

type AuditLog struct {
	ID          uint `gorm:"primaryKey"`
	UserID      *uint
	User        *User
	TenantID    uint
	Tenant      Tenant
	Module      string
	Action      string
	Level       string
	ReferenceID string
	Message     string
	DataBefore  datatypes.JSON
	DataAfter   datatypes.JSON
	IPAddress   string
	UserAgent   string
	CreatedAt   time.Time
}

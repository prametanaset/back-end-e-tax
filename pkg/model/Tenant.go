package model

type Tenant struct {
	ID         uint   `gorm:"primaryKey"`
	LegalName  string `gorm:"not null"`
	BranchCode string
	TaxID      string
	AddressID  uint
	Address    Address `gorm:"foreignKey:AddressID"`
	LogoURL    string
	Users      []User `gorm:"foreignKey:TenantID"`
}

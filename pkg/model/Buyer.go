package model

type Buyer struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	TaxID         string
	AddressID     uint
	Address       Address
	ContactPerson string
	Email         string
}

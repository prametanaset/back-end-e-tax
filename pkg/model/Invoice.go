package model

import "time"

type Invoice struct {
	ID                uint `gorm:"primaryKey"`
	TenantID          uint
	Tenant            Tenant
	BuyerID           uint
	Buyer             Buyer
	InvoiceNumber     string `gorm:"uniqueIndex:tenant_invoice_number"`
	IssueDate         time.Time
	CurrencyCode      string `gorm:"default:THB"`
	SubTotal          float64
	DiscountGlobalPct float64
	DiscountGlobalAmt float64
	FreightTotal      float64
	VatBaseTotal      float64
	VatTotal          float64
	GrandTotal        float64
	Status            string
	XMLPath           string
	Items             []InvoiceItem
	Payments          []Payment
}

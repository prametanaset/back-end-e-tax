package model

type InvoiceItem struct {
	ID              uint `gorm:"primaryKey"`
	InvoiceID       uint
	Invoice         Invoice
	ProductID       uint
	Product         Product
	Description     string
	SeqNo           int
	Qty             float64
	UOM             string
	UnitPrice       float64
	LineDiscountPct float64
	LineDiscountAmt float64
	VatRate         float64
	VatBase         float64
	VatAmt          float64
	LineTotal       float64
}

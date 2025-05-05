package model

import "time"

type Payment struct {
	ID         uint `gorm:"primaryKey"`
	InvoiceID  uint
	Invoice    Invoice
	Method     string
	PaidAmount float64
	PaidDate   time.Time
	RefNo      string
	Reconcile  bool
}

package migration

import (
	"back-end-e-tax/pkg/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Address{},
		&model.Tenant{},
		&model.Buyer{},
		&model.Invoice{},
		&model.InvoiceItem{},
		&model.Product{},
		&model.VatCategory{},
		&model.TaxRate{},
		&model.Payment{},
		&model.User{},
		&model.AuthToken{},
		&model.Role{},
		&model.AuditLog{},
	)
}

// pkg/repository/tenant_repository.go
package repository

import (
	"back-end-e-tax/pkg/model"

	"gorm.io/gorm"
)

type TenantRepository interface {
	Create(t *model.Tenant) error
}

type tenantRepository struct{ db *gorm.DB }

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db}
}

func (r *tenantRepository) Create(t *model.Tenant) error {
	return r.db.Create(t).Error
}

package repository

import (
	"back-end-e-tax/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(tenantID uint) ([]model.Product, error)
	GetByID(id int, tenantID uint) (*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(product *model.Product) error
	GetLastProductCode(tenantID uint) (*model.Product, error)
	GetPaginated(page, limit int, search string, tenantID uint) ([]model.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll(tenantID uint) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("tenant_id = ?", tenantID).Order("updated_at DESC").Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id int, tenantID uint) (*model.Product, error) {
	var product model.Product
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&product).Error
	return &product, err
}

func (r *productRepository) GetPaginated(page, limit int, search string, tenantID uint) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64
	offset := (page - 1) * limit

	query := r.db.Model(&model.Product{}).Where("tenant_id = ?", tenantID)

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("name LIKE ? OR product_code LIKE ?", searchPattern, searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("updated_at DESC").Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(product *model.Product) error {
	return r.db.Delete(product).Error
}

func (r *productRepository) GetLastProductCode(tenantID uint) (*model.Product, error) {
	var product model.Product
	err := r.db.Where("tenant_id = ?", tenantID).Order("id DESC").First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &product, err
}

package repository

import (
	"prodexa/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll(tenantID uint) ([]models.Product, error)
	GetByID(id, tenantID uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id, tenantID uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

// CREATE
func (r *productRepo) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// READ ALL
func (r *productRepo) GetAll(tenantID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("tenant_id = ?", tenantID).Find(&products).Error
	return products, err
}

// READ ONE (WITH TENANT CHECK ❗)
func (r *productRepo) GetByID(id, tenantID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&product).Error
	return &product, err
}

// UPDATE (SAFE)
func (r *productRepo) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// DELETE (WITH TENANT CHECK ❗)
func (r *productRepo) Delete(id, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&models.Product{}).Error
}

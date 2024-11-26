package gorm

import (
	model "go-api-test/internal/models"

	"gorm.io/gorm"
)

type GormProductsRepository struct {
	db *gorm.DB
}

func NewGormProductsRepository(db *gorm.DB) *GormProductsRepository {
	return &GormProductsRepository{db: db}
}

func (r *GormProductsRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *GormProductsRepository) FindByID(id string) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, "id = ?", id).Error
	return &product, err
}

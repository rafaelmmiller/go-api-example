package repository

import (
	"go-api-test/model"
)

type ProductsRepository interface {
	// Create(product *model.Product) error
	FindByID(id string) (*model.Product, error)
	GetAll() ([]model.Product, error)
}

// func (r *GormProductRepository) Create(product *model.Product) error {
// 	return r.db.Create(product).Error
// }

package usecases

import (
	model "go-api-test/internal/models"
	repository "go-api-test/internal/repository"
)

type CreateProductUseCase struct {
	productRepo repository.ProductsRepository
}

func NewCreateProductUseCase(productRepo repository.ProductsRepository) *CreateProductUseCase {
	return &CreateProductUseCase{productRepo: productRepo}
}

func (uc *CreateProductUseCase) CreateProduct(product *model.Product) (*model.Product, error) {
	return nil, nil // uc.productRepo.Create(product)
}

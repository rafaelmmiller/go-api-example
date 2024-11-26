package usecases

import (
	model "go-api-test/internal/models"
	repository "go-api-test/internal/repository"
)

type GetByIdProductUseCase struct {
	productRepo repository.ProductsRepository
}

func NewGetByIdProductUseCase(productRepo repository.ProductsRepository) *GetByIdProductUseCase {
	return &GetByIdProductUseCase{productRepo: productRepo}
}

func (uc *GetByIdProductUseCase) GetByIdProducts(id string) (*model.Product, error) {
	return uc.productRepo.FindByID(id)
}

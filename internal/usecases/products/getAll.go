package usecases

import (
	model "go-api-test/internal/models"
	repository "go-api-test/internal/repository"
)

type GetAllProductUseCase struct {
	productRepo repository.ProductsRepository
}

func NewGetAllProductUseCase(productRepo repository.ProductsRepository) *GetAllProductUseCase {
	return &GetAllProductUseCase{productRepo: productRepo}
}

func (uc *GetAllProductUseCase) GetAllProducts() ([]model.Product, error) {
	return uc.productRepo.GetAll()
}

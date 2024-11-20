package usecases

import (
	"go-api-test/model"
	"go-api-test/repository"
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

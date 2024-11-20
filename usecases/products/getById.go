package usecases

import (
	"go-api-test/model"
	"go-api-test/repository"
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

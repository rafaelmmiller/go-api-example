package controller

import (
	usecases "go-api-test/usecases/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProductsController struct {
	getAllProductUseCase *usecases.GetAllProductUseCase
}

func NewGetProductsController(getAllProductUseCase *usecases.GetAllProductUseCase) *GetProductsController {
	return &GetProductsController{getAllProductUseCase: getAllProductUseCase}
}

func (pc *GetProductsController) GetAllProducts(ctx *gin.Context) {
	products, err := pc.getAllProductUseCase.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

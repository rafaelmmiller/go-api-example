package controller

import (
	usecases "go-api-test/internal/usecases/products"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetByIdProductsController struct {
	getByIdProductUseCase *usecases.GetByIdProductUseCase
}

func NewGetByIdProductsController(getByIdProductUseCase *usecases.GetByIdProductUseCase) *GetByIdProductsController {
	return &GetByIdProductsController{getByIdProductUseCase: getByIdProductUseCase}
}

func (pc *GetByIdProductsController) GetByIdProducts(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	product, err := pc.getByIdProductUseCase.GetByIdProducts(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

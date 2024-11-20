package controller

import (
	"fmt"
	"go-api-test/model"
	usecases "go-api-test/usecases/products"
	"go-api-test/utils"
	"net/http"

	z "github.com/Oudwins/zog"

	"github.com/gin-gonic/gin"
)

type CreateProductsController struct {
	createProductUseCase *usecases.CreateProductUseCase
}

func NewCreateProductsController(createProductUseCase *usecases.CreateProductUseCase) *CreateProductsController {
	return &CreateProductsController{createProductUseCase: createProductUseCase}
}

func (pc *CreateProductsController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	jsonBody, err := utils.UnmarshalRequestBody(ctx.Request.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	zErr := model.ProductSchema.Parse(jsonBody, &product)
	if zErr != nil {
		sanitized := z.Errors.SanitizeMap(zErr)
		delete(sanitized, "$first") // idk why this returns $first...
		ctx.JSON(http.StatusBadRequest, gin.H{"error": sanitized})
		return
	}

	productPtr, err := pc.createProductUseCase.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    productPtr,
	})
}

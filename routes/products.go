package routes

import (
	"github.com/gin-gonic/gin"

	controller "go-api-test/controllers/products"
	db "go-api-test/db"
	gorm "go-api-test/repository/implementations/gorm"
	usecases "go-api-test/usecases/products"
)

func RegisterRoutes(router *gin.Engine) {
	registerProductRoutes(router)
}

func registerProductRoutes(router *gin.Engine) {
	db := db.Init()
	productRepo := gorm.NewGormProductsRepository(db)

	// Get All Products
	getAllProductsUseCase := usecases.NewGetAllProductUseCase(productRepo)
	getProductsController := controller.NewGetProductsController(getAllProductsUseCase)
	router.GET("/products", getProductsController.GetAllProducts)

	// Get Product By Id
	getByIdProductsUseCase := usecases.NewGetByIdProductUseCase(productRepo)
	getByIdProductsController := controller.NewGetByIdProductsController(getByIdProductsUseCase)
	router.GET("/products/:id", getByIdProductsController.GetByIdProducts)

	// Create Product
	createProductUseCase := usecases.NewCreateProductUseCase(productRepo)
	createProductController := controller.NewCreateProductsController(createProductUseCase)
	router.POST("/products", createProductController.CreateProduct)
}

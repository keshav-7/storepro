package main

import (
	"prodexa/internal/database"
	"prodexa/internal/handler"
	"prodexa/internal/repository"
	"prodexa/internal/routes"
	"prodexa/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// DB
	db := database.InitDB()

	// Dependency Injection
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Routes
	routes.SetupRoutes(r, productHandler)

	// Start server
	r.Run(":8080")
}

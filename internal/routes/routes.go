package routes

import (
	"storepro/internal/handler"
	"storepro/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, productHandler *handler.ProductHandler) {

	api := r.Group("/api")

	// Public routes
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	{
		products := protected.Group("/products")
		{
			products.POST("/", productHandler.CreateProduct)
			products.GET("/", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
	}
}

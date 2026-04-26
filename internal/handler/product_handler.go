package handler

import (
	"log"
	"net/http"
	"strconv"

	"storepro/internal/models"
	"storepro/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

// CREATE
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Debug log
	log.Println("Incoming product:", product)

	if product.Name == "" {
		c.JSON(400, gin.H{"error": "name is required"})
		return
	}

	tenantID := c.GetUint("tenant_id")

	err := h.service.CreateProduct(&product, tenantID)
	if err != nil {
		log.Println("Create error:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, product)
}

// GET ALL
func (h *ProductHandler) GetProducts(c *gin.Context) {
	tenantID := c.GetUint("tenant_id")

	products, err := h.service.GetProducts(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GET ONE
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tenantID := c.GetUint("tenant_id")

	product, err := h.service.GetProduct(uint(id), tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UPDATE
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tenantID := c.GetUint("tenant_id")

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)

	if err := h.service.UpdateProduct(&product, tenantID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DELETE
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tenantID := c.GetUint("tenant_id")

	if err := h.service.DeleteProduct(uint(id), tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

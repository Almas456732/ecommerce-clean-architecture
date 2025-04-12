package http

import (
	"net/http"
	"strconv"

	"inventory_service/internal/application"
	"inventory_service/internal/domain"
	"inventory_service/internal/dto"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *application.ProductService
}

type CategoryHandler struct {
	service *application.CategoryService
}

func SetupRoutes(r *gin.Engine, productService *application.ProductService, categoryService *application.CategoryService) {
	productHandler := &ProductHandler{service: productService}
	categoryHandler := &CategoryHandler{service: categoryService}

	api := r.Group("/inventory")
	{

		products := api.Group("/products")
		{
			products.POST("/", productHandler.CreateProduct)
			products.GET("/:id", productHandler.GetProduct)
			products.PATCH("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.GET("/", productHandler.ListProducts)
		}

		categories := api.Group("/categories")
		{
			categories.POST("/", categoryHandler.CreateCategory)
			categories.GET("/:id", categoryHandler.GetCategory)
			categories.PATCH("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
			categories.GET("/", categoryHandler.ListCategories)
		}
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := productDTO.ToDomain()
	if err := h.service.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productDTO.FromDomain(product)
	c.JSON(http.StatusCreated, productDTO)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.service.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var productDTO dto.ProductDTO
	productDTO.FromDomain(product)
	c.JSON(http.StatusOK, productDTO)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var productDTO dto.ProductDTO
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productDTO.ID = id
	product := productDTO.ToDomain()

	if err := h.service.UpdateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productDTO.FromDomain(product)
	c.JSON(http.StatusOK, productDTO)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filter := domain.ProductFilter{
		Name:       c.Query("name"),
		CategoryID: c.Query("category_id"),
	}

	products, err := h.service.ListProducts(filter, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productDTOs := make([]dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i].FromDomain(product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var categoryDTO dto.CategoryDTO
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := categoryDTO.ToDomain()
	if err := h.service.CreateCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryDTO.FromDomain(category)
	c.JSON(http.StatusCreated, categoryDTO)
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := h.service.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	var categoryDTO dto.CategoryDTO
	categoryDTO.FromDomain(category)
	c.JSON(http.StatusOK, categoryDTO)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var categoryDTO dto.CategoryDTO
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryDTO.ID = id
	category := categoryDTO.ToDomain()

	if err := h.service.UpdateCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryDTO.FromDomain(category)
	c.JSON(http.StatusOK, categoryDTO)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.service.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoryDTOs := make([]dto.CategoryDTO, len(categories))
	for i, category := range categories {
		categoryDTOs[i].FromDomain(category)
	}

	c.JSON(http.StatusOK, categoryDTOs)
}

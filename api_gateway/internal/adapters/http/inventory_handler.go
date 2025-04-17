package http

import (
	"net/http"
	"strconv"

	"api_gateway/internal/adapters/grpc"
	pb "api_gateway/proto/inventory"

	"github.com/gin-gonic/gin"
)

// InventoryHandler handles HTTP requests for inventory operations
type InventoryHandler struct {
	inventoryClient *grpc.InventoryServiceClient
}

// NewInventoryHandler creates a new instance of InventoryHandler
func NewInventoryHandler(inventoryClient *grpc.InventoryServiceClient) *InventoryHandler {
	return &InventoryHandler{
		inventoryClient: inventoryClient,
	}
}

// SetupRoutes configures the routes for the inventory service
func (h *InventoryHandler) SetupRoutes(router *gin.Engine) {
	inventory := router.Group("/inventory")
	{
		// Product routes
		products := inventory.Group("/products")
		{
			products.GET("", h.ListProducts)
			products.GET("/:id", h.GetProduct)
			products.POST("", h.CreateProduct)
			products.PUT("/:id", h.UpdateProduct)
			products.DELETE("/:id", h.DeleteProduct)
		}

		// Category routes
		categories := inventory.Group("/categories")
		{
			categories.GET("", h.ListCategories)
			categories.GET("/:id", h.GetCategory)
			categories.POST("", h.CreateCategory)
			categories.PUT("/:id", h.UpdateCategory)
			categories.DELETE("/:id", h.DeleteCategory)
		}
	}
}

// CreateProduct creates a new product
func (h *InventoryHandler) CreateProduct(c *gin.Context) {
	var request struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		CategoryID  string  `json:"category_id" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		Stock       int     `json:"stock" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.inventoryClient.CreateProduct(c.Request.Context(), &pb.CreateProductRequest{
		Name:        request.Name,
		Description: request.Description,
		CategoryId:  request.CategoryID,
		Price:       request.Price,
		Stock:       int32(request.Stock),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProduct gets a product by ID
func (h *InventoryHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.inventoryClient.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates a product
func (h *InventoryHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		CategoryID  string  `json:"category_id" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		Stock       int     `json:"stock" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.inventoryClient.UpdateProduct(c.Request.Context(), &pb.UpdateProductRequest{
		Id:          id,
		Name:        request.Name,
		Description: request.Description,
		CategoryId:  request.CategoryID,
		Price:       request.Price,
		Stock:       int32(request.Stock),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product
func (h *InventoryHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.inventoryClient.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// ListProducts lists all products
func (h *InventoryHandler) ListProducts(c *gin.Context) {
	// Parse query parameters
	name := c.Query("name")
	categoryID := c.Query("category_id")

	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("min_price", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("max_price", "0"), 64)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filter := &pb.ProductFilter{
		Name:       name,
		CategoryId: categoryID,
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
	}

	response, err := h.inventoryClient.ListProducts(c.Request.Context(), filter, int32(page), int32(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list products: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Products)
}

// CreateCategory creates a new category
func (h *InventoryHandler) CreateCategory(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.inventoryClient.CreateCategory(c.Request.Context(), &pb.CreateCategoryRequest{
		Name:        request.Name,
		Description: request.Description,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategory gets a category by ID
func (h *InventoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := h.inventoryClient.GetCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates a category
func (h *InventoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.inventoryClient.UpdateCategory(c.Request.Context(), &pb.UpdateCategoryRequest{
		Id:          id,
		Name:        request.Name,
		Description: request.Description,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory deletes a category
func (h *InventoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := h.inventoryClient.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// ListCategories lists all categories
func (h *InventoryHandler) ListCategories(c *gin.Context) {
	response, err := h.inventoryClient.ListCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list categories: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Categories)
}

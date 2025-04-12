package http

import (
	"net/http"

	"order_service/internal/application"
	"order_service/internal/dto"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *application.OrderService
}

func SetupRoutes(r *gin.Engine, orderService *application.OrderService) {
	handler := &OrderHandler{service: orderService}

	api := r.Group("/orders")
	{

		orders := api.Group("/orders")
		{
			orders.POST("/", handler.CreateOrder)
			orders.GET("/:id", handler.GetOrder)
			orders.PATCH("/:id", handler.UpdateOrderStatus)
			orders.GET("/", handler.GetUserOrders)
		}
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderDTO dto.OrderDTO
	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := orderDTO.ToDomain()
	if err := h.service.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderDTO.FromDomain(order)
	c.JSON(http.StatusCreated, orderDTO)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	var orderDTO dto.OrderDTO
	orderDTO.FromDomain(order)
	c.JSON(http.StatusOK, orderDTO)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var orderDTO dto.OrderDTO
	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderDTO.ID = id

	order := orderDTO.ToDomain()
	if err := h.service.UpdateOrderStatus(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderDTO.FromDomain(order)
	c.JSON(http.StatusOK, orderDTO)
}

func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	orders, err := h.service.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderDTOs := make([]dto.OrderDTO, len(orders))
	for i, order := range orders {
		orderDTOs[i].FromDomain(order)
	}

	c.JSON(http.StatusOK, orderDTOs)
}

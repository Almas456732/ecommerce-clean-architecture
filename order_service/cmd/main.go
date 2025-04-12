package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"order_service/internal/adapters/http"
	"order_service/internal/adapters/repositories"
	"order_service/internal/application"
	"order_service/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	
	orderRepo := repositories.NewOrderRepository(cfg.DatabaseURL)

	
	orderService := application.NewOrderService(orderRepo)

	
	r := gin.Default()
	http.SetupRoutes(r, orderService)

	
	go func() {
		log.Printf("Starting order service on port %s", cfg.Port)
		if err := r.Run(":" + cfg.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

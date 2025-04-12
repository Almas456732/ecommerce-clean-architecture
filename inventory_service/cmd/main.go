package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"inventory_service/internal/adapters/http"
	"inventory_service/internal/adapters/repositories"
	"inventory_service/internal/application"
	"inventory_service/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	
	productRepo := repositories.NewProductRepository(cfg.DatabaseURL)
	categoryRepo := repositories.NewCategoryRepository(cfg.DatabaseURL)

	
	productService := application.NewProductService(productRepo)
	categoryService := application.NewCategoryService(categoryRepo)

	
	r := gin.Default()
	http.SetupRoutes(r, productService, categoryService)

	
	go func() {
		log.Printf("Starting inventory service on port %s", cfg.Port)
		if err := r.Run(":" + cfg.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}

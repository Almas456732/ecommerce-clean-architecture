package main

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"api_gateway/internal/adapters/http"
	"api_gateway/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	inventoryProxy := createReverseProxy(cfg.InventoryServiceURL)
	orderProxy := createReverseProxy(cfg.OrderServiceURL)
	userProxy := createReverseProxy(cfg.UserServiceURL)

	router := gin.Default()

	router.Use(http.LoggingMiddleware())
	router.Use(http.AuthMiddleware(cfg.JWTSecret))

	http.SetupRoutes(router, inventoryProxy, orderProxy, userProxy)

	go func() {
		log.Printf("API Gateway running on port %s", cfg.Port)
		if err := router.Run(":" + cfg.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down API Gateway...")
}

func createReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

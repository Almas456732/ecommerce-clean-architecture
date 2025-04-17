package main

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	grpcClient "api_gateway/internal/adapters/grpc"
	"api_gateway/internal/adapters/http"
	"api_gateway/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize gRPC clients
	inventoryClient, err := grpcClient.NewInventoryServiceClient(cfg.InventoryServiceGRPCURL)
	if err != nil {
		log.Fatalf("Failed to connect to inventory gRPC service: %v", err)
	}
	defer inventoryClient.Close()

	// Create HTTP handlers
	inventoryHandler := http.NewInventoryHandler(inventoryClient)

	// Create legacy reverse proxies for services not yet migrated to gRPC
	orderProxy := createReverseProxy(cfg.OrderServiceURL)
	userProxy := createReverseProxy(cfg.UserServiceURL)

	router := gin.Default()

	router.Use(http.LoggingMiddleware())
	router.Use(http.AuthMiddleware(cfg.JWTSecret))

	// Set up routes for inventory service using gRPC client
	inventoryHandler.SetupRoutes(router)

	// Set up routes for other services still using proxies
	http.SetupProxyRoutes(router, orderProxy, userProxy)

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

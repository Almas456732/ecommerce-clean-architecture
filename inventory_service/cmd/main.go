package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"inventory_service/internal/adapters/grpc"
	"inventory_service/internal/adapters/http"
	"inventory_service/internal/adapters/repositories"
	"inventory_service/internal/application"
	"inventory_service/internal/config"
	pb "inventory_service/proto/inventory"

	"github.com/gin-gonic/gin"
	grpcserver "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize repositories
	productRepo := repositories.NewProductRepository(cfg.DatabaseURL)
	categoryRepo := repositories.NewCategoryRepository(cfg.DatabaseURL)

	// Initialize services
	productService := application.NewProductService(productRepo)
	categoryService := application.NewCategoryService(categoryRepo)

	// Start HTTP server
	r := gin.Default()
	http.SetupRoutes(r, productService, categoryService)

	go func() {
		log.Printf("Starting HTTP server on port %s", cfg.Port)
		if err := r.Run(":" + cfg.Port); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}

	grpcServer := grpcserver.NewServer()
	server := grpc.NewServer(productService, categoryService)

	pb.RegisterProductServiceServer(grpcServer, server)
	pb.RegisterCategoryServiceServer(grpcServer, server)

	// Enable reflection for gRPC tools
	reflection.Register(grpcServer)

	go func() {
		log.Printf("Starting gRPC server on port %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")
	grpcServer.GracefulStop()
}

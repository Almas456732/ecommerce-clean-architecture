package main

import (
	"fmt"
	"log"

	"user_service/internal/adapters/auth"
	"user_service/internal/adapters/http"
	"user_service/internal/adapters/repositories"
	"user_service/internal/application"
	"user_service/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create repository
	userRepo, err := repositories.NewMongoUserRepository(
		cfg.MongoURI,
		cfg.MongoDatabase,
		cfg.MongoUserCollection,
	)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create services
	userService := application.NewUserService(userRepo)
	tokenService := auth.NewTokenService(cfg.JWTSecret)

	// Create HTTP handler
	userHandler := http.NewUserHandler(userService, tokenService)

	// Setup router
	router := gin.Default()

	// Setup routes
	http.SetupRoutes(router, userHandler)

	// Start the server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting user service on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

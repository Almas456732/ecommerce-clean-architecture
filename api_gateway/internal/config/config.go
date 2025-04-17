package config

import "os"

type Config struct {
	Port                    string
	InventoryServiceURL     string
	InventoryServiceGRPCURL string
	OrderServiceURL         string
	OrderServiceGRPCURL     string
	UserServiceURL          string
	UserServiceGRPCURL      string
	JWTSecret               string
}

func LoadConfig() *Config {
	return &Config{
		Port:                    getEnv("PORT", "8080"),
		InventoryServiceURL:     getEnv("INVENTORY_SERVICE_URL", "http://localhost:8081"),
		InventoryServiceGRPCURL: getEnv("INVENTORY_SERVICE_GRPC_URL", "localhost:50051"),
		OrderServiceURL:         getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
		OrderServiceGRPCURL:     getEnv("ORDER_SERVICE_GRPC_URL", "localhost:50052"),
		UserServiceURL:          getEnv("USER_SERVICE_URL", "http://localhost:8083"),
		UserServiceGRPCURL:      getEnv("USER_SERVICE_GRPC_URL", "localhost:50053"),
		JWTSecret:               getEnv("JWT_SECRET", "super-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

package config

import "os"

type Config struct {
	Port                string
	InventoryServiceURL string
	OrderServiceURL     string
	JWTSecret           string
}

func LoadConfig() *Config {
	return &Config{
		Port:                getEnv("PORT", "8080"),
		InventoryServiceURL: getEnv("INVENTORY_SERVICE_URL", "http://localhost:8081"),
		OrderServiceURL:     getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
		JWTSecret:           getEnv("JWT_SECRET", "super-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

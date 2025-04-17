package config

import "os"

type Config struct {
	Port        string
	GRPCPort    string
	DatabaseURL string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8081"),
		GRPCPort:    getEnv("GRPC_PORT", "50051"),
		DatabaseURL: getEnv("DATABASE_URL", "mongodb://localhost:27017"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

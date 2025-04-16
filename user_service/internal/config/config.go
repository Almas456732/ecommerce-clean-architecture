package config

import "os"

type Config struct {
	Port                string
	JWTSecret           string
	MongoURI            string
	MongoDatabase       string
	MongoUserCollection string
}

func LoadConfig() *Config {
	return &Config{
		Port:                getEnv("PORT", "8083"),
		JWTSecret:           getEnv("JWT_SECRET", "super-secret-key"),
		MongoURI:            getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDatabase:       getEnv("MONGO_DATABASE", "user_service"),
		MongoUserCollection: getEnv("MONGO_USER_COLLECTION", "users"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

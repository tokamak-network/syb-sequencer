package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	// Database configuration
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Ethereum configuration
	EthereumRPC     string
	ContractAddress string

	// Synchronizer configuration
	PollingInterval     int // in seconds
	SafetyCheckInterval int // in seconds
	MaxBlocksPerBatch   int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	config := &Config{
		// Database configuration
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "sybil"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		// Ethereum configuration
		EthereumRPC:     getEnv("ETHEREUM_RPC", "http://localhost:8545"),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),

		// Synchronizer configuration with defaults
		PollingInterval:     getEnvInt("POLLING_INTERVAL", 15),
		SafetyCheckInterval: getEnvInt("SAFETY_CHECK_INTERVAL", 300), // 5 minutes
		MaxBlocksPerBatch:   getEnvInt("MAX_BLOCKS_PER_BATCH", 1000),
	}

	return config
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvInt gets an environment variable as an integer or returns a default value
func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

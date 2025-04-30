package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	// Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	// Blockchain
	EthereumRPC     string
	ContractAddress string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "sequencer"),

		// Blockchain
		EthereumRPC:     getEnv("ETHEREUM_RPC", "https://ethereum-sepolia-rpc.publicnode.com"),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

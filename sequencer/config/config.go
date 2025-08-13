package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	// HistoryDB configuration
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// StateDB configuration
	Path string
	Keep int

	// Ethereum configuration
	EthereumRPC     string
	ContractAddress string

	// Synchronizer configuration
	PollingInterval     int // in seconds
	SafetyCheckInterval int // in seconds
	MaxBlocksPerBatch   int
}

type DBCredentials struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	appMode := GetEnv("APP_MODE", "dev")

	if appMode == "test" {
		var secrets string
		var err error

		secrets, err = getAwsSecrets(os.Getenv("AWS_DB_SECRET_NAME"), os.Getenv("AWS_REGION"))
		if secrets == "" {
			panic("No secrets found in AWS Secrets Manager.")
		}
		var creds DBCredentials
		err = json.Unmarshal([]byte(secrets), &creds)
		if err != nil {
			panic(err)
		}
		// TODO: for some reason RDS doesn't save dbname in the secrets, need to investigate
		if creds.Dbname == "" {
			creds.Dbname = "postgres"
		}

		config := &Config{
			// HistoryDB configuration
			DBHost:     creds.Host,
			DBPort:     creds.Port,
			DBUser:     creds.Username,
			DBPassword: creds.Password,
			DBName:     creds.Dbname,
			DBSSLMode:  "require",

			// StateDB configuration
			Path: GetEnv("STATEDB_DIR_PATH", "./var/tokamak/statedb"),
			Keep: GetEnvInt("KEEP", 256),

			// Ethereum configuration
			EthereumRPC:     GetEnv("ETHEREUM_RPC", "http://localhost:8545"),
			ContractAddress: GetEnv("CONTRACT_ADDRESS", ""),

			// Synchronizer configuration with defaults
			PollingInterval:     GetEnvInt("POLLING_INTERVAL", 15),
			SafetyCheckInterval: GetEnvInt("SAFETY_CHECK_INTERVAL", 300), // 5 minutes
			MaxBlocksPerBatch:   GetEnvInt("MAX_BLOCKS_PER_BATCH", 1000),
		}

		return config
	}

	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	config := &Config{
		// HistoryDB configuration
		DBHost:     GetEnv("DB_HOST", "localhost"),
		DBPort:     GetEnvInt("DB_PORT", 5432),
		DBUser:     GetEnv("DB_USER", "postgres"),
		DBPassword: GetEnv("DB_PASSWORD", "postgres"),
		DBName:     GetEnv("DB_NAME", "sybil"),
		DBSSLMode:  GetEnv("DB_SSLMODE", "disable"),

		// StateDB configuration
		Path: GetEnv("STATEDB_DIR_PATH", "./var/tokamak/statedb"),
		Keep: GetEnvInt("KEEP", 256),

		// Ethereum configuration
		EthereumRPC:     GetEnv("ETHEREUM_RPC", "http://localhost:8545"),
		ContractAddress: GetEnv("CONTRACT_ADDRESS", ""),

		// Synchronizer configuration with defaults
		PollingInterval:     GetEnvInt("POLLING_INTERVAL", 15),
		SafetyCheckInterval: GetEnvInt("SAFETY_CHECK_INTERVAL", 300), // 5 minutes
		MaxBlocksPerBatch:   GetEnvInt("MAX_BLOCKS_PER_BATCH", 1000),
	}

	return config
}

// getEnv gets an environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvInt gets an environment variable as an integer or returns a default value
func GetEnvInt(key string, defaultValue int) int {
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

// getEnvInt64 gets an environment variable as an int64 or returns a default value
func GetEnvInt64(key string, defaultValue int64) int64 {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func getAwsSecrets(secretName, region string) (string, error) {
	if secretName == "" || region == "" {
		return "", fmt.Errorf("secretName and region must be provided to fetch secrets from AWS")
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
		Profile:           "default",
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return "", err
	}
	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}
	return *result.SecretString, nil
}

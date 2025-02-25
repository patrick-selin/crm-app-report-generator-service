package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	S3Bucket     string
	DynamoDBTable string
	AWSRegion    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	return &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DBHost:       getEnv("POSTGRES_HOST", "localhost"),
		DBPort:       getEnv("POSTGRES_PORT", "5432"),
		DBUser:       getEnv("POSTGRES_USER", "patrick"),
		DBPassword:   getEnv("POSTGRES_PASSWORD", ""),
		DBName:       getEnv("POSTGRES_DB", "crm_app_db_dev"),
		S3Bucket:     getEnv("S3_BUCKET", "report-generator-bucket"),
		DynamoDBTable: getEnv("DYNAMODB_TABLE", "ReportMetadataTable"),
		AWSRegion:    getEnv("AWS_REGION", "eu-north-1"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

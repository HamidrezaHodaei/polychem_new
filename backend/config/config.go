package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT
	JWTSecret              string
	JWTExpiryHours         int
	RefreshTokenExpiryDays int

	// Server
	ServerPort  string
	FrontendURL string

	// Security
	BcryptCost int

	Environment string
}

var AppConfig *Config

func Load() {
	// بارگذاری فایل .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	AppConfig = &Config{
		DBHost:                 getEnv("DB_HOST", "localhost"),
		DBPort:                 getEnv("DB_PORT", "3306"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBName:                 getEnv("DB_NAME", "polychem_db"),
		JWTSecret:              getEnv("JWT_SECRET", "change-this-secret-key"),
		JWTExpiryHours:         getEnvInt("JWT_EXPIRY_HOURS", 24),
		RefreshTokenExpiryDays: getEnvInt("REFRESH_TOKEN_EXPIRY_DAYS", 7),
		ServerPort:             getEnv("SERVER_PORT", "8080"),
		FrontendURL:            getEnv("FRONTEND_URL", "http://localhost:3000"),
		BcryptCost:             getEnvInt("BCRYPT_COST", 12),
		Environment:            getEnv("ENVIRONMENT", "development"),
	}

	validateConfig()
}

func validateConfig() {
	if AppConfig.JWTSecret == "change-this-secret-key" {
		log.Fatal("FATAL: JWT_SECRET must be changed in production!")
	}
	if AppConfig.DBPassword == "" {
		log.Println("WARNING: Database password is empty")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
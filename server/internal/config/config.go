// Package config contains a configuration loader for environment variables
package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port          string
	DBHost        string
	DBPort        string
	Domain        string
	DBUser        string
	DBPassword    string
	GoEnv         string
	DBName        string
	JWTSecret     string
	RefreshSecret string
}

// LoadConfig reads environment variables and returns a Config struct
func LoadConfig() *Config {
	return &Config{
		Port:          getEnv("PORT", ""),
		Domain:        getEnv("DOMAIN", ""),
		DBHost:        getEnv("DB_HOST", ""),
		DBPort:        getEnv("DB_PORT", ""),
		GoEnv:         getEnv("GO_ENV", ""),
		DBUser:        getEnv("DB_USER", ""),
		DBPassword:    getEnv("DB_PASSWORD", ""),
		DBName:        getEnv("DB_NAME", ""),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		RefreshSecret: getEnv("REFRESH_SECRET", ""),
	}
}

// getEnv reads an env variable or returns default
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

//	func ConstructDBString(c *Config) string {
//		return fmt.Sprintf(
//			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
//			c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
//		)
//	}
func ConstructDBString(cfg Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=require",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBName,
	)
}

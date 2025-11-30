
package config

import (
    "fmt"
    "os"
)

type Config struct {
    Port       string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
}

// LoadConfig reads environment variables and returns a Config struct
func LoadConfig() *Config {
    return &Config{
        Port:       getEnv("PORT", "8080"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "paste_store"),
        JWTSecret:  getEnv("JWT_SECRET", "supersecret"),
    }
}

// getEnv reads an env variable or returns default
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// Helper to construct a PostgreSQL DSN
func (c *Config) PostgresDSN() string {
    return fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
    )
}

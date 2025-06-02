package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, reading from environment variables")
    }

    cfg := Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "3306"),
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", "root"),
        DBName:     getEnv("DB_NAME", "go_assignment"),
    }

    return cfg
}

func getEnv(key string, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

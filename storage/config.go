package storage

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host    string `env:"DB_HOST"`
	Port    string `env:"DB_PORT"`
	User    string `env:"DB_USER"`
	Pass    string `env:"DB_PASSWORD"`
	Name    string `env:"DB_NAME"`
	SSLMode string `env:"DB_SSL_MODE"`
}

func NewStorageConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config

	cfg.Host = getEnvOfFatal("DB_HOST")
	cfg.Port = getEnvOfFatal("DB_PORT")
	cfg.User = getEnvOfFatal("DB_USER")
	cfg.Pass = getEnvOfFatal("DB_PASSWORD")
	cfg.Name = getEnvOfFatal("DB_NAME")
	cfg.SSLMode = getEnvOfFatal("DB_SSL_MODE")

	log.Println("config successfully")

	return &cfg

}

func getEnvOfFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable not set", key)
	}

	return value
}

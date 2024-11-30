package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Host    string `env:"DB_HOST" env-Default:"localhost"`
	Port    string `env:"DB_PORT" env-Default:"5435"`
	User    string `env:"DB_USER" env-Default:"postgres"`
	Pass    string `env:"DB_PASSWORD" env-Default:"postgres"`
	Name    string `env:"DB_NAME" env-Default:"postgres"`
	SSLMode string `env:"DB_SSL_MODE" env-Default:"disable"`
}

func NewStorageConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Println("the configuration with default values has been loaded")

		return nil
	}

	log.Println("config successfully")

	return &cfg

}

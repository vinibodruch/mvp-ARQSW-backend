package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	OmdbAPIKey  string
	Port        string
}

func Load() *Config {
	// Carrega .env se existir (ignorado em produção Docker)
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL não definida")
	}

	omdbKey := os.Getenv("OMDB_API_KEY")
	if omdbKey == "" {
		log.Fatal("OMDB_API_KEY não definida")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		OmdbAPIKey:  omdbKey,
		Port:        port,
	}
}

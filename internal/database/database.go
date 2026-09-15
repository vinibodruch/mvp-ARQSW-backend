package database

import (
	"log"

	"movie-watchlist-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar no banco: %v", err)
	}

	// Cria a tabela automaticamente se não existir
	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatalf("Falha ao migrar banco: %v", err)
	}

	log.Println("Banco de dados conectado e migrado com sucesso")
	return db
}

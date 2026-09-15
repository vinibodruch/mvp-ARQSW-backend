package main

import (
	"log"

	"movie-watchlist-api/internal/config"
	"movie-watchlist-api/internal/database"
	"movie-watchlist-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)

	router := gin.Default()

	// Permite requisições do front-end
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	movieHandler := handlers.NewMovieHandler(db)

	api := router.Group("/api")
	{
		api.GET("/search", handlers.SearchMovie(cfg.OmdbAPIKey))
		api.GET("/movies", movieHandler.ListMovies)
		api.POST("/movies", movieHandler.CreateMovie)
		api.PUT("/movies/:id", movieHandler.UpdateMovie)
		api.DELETE("/movies/:id", movieHandler.DeleteMovie)
	}

	log.Printf("Servidor rodando na porta %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}

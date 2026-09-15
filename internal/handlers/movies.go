package handlers

import (
	"net/http"

	"movie-watchlist-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieHandler struct {
	DB *gorm.DB
}

func NewMovieHandler(db *gorm.DB) *MovieHandler {
	return &MovieHandler{DB: db}
}

// ListMovies retorna todos os filmes da watchlist, com filtro opcional por watched
func (h *MovieHandler) ListMovies(c *gin.Context) {
	var movies []models.Movie
	query := h.DB

	// Filtro opcional: /api/movies?watched=true
	if watched := c.Query("watched"); watched != "" {
		isWatched := watched == "true"
		query = query.Where("is_watched = ?", isWatched)
	}

	if err := query.Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar filmes"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// CreateMovie adiciona um filme à watchlist.
// Se o filme foi deletado anteriormente (soft delete), restaura o registro em vez de inserir.
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Busca registro existente incluindo soft-deleted
	var existing models.Movie
	err := h.DB.Unscoped().Where("imdb_id = ?", input.ImdbID).First(&existing).Error

	if err == nil {
		if existing.DeletedAt.Valid {
			// Restaura o registro deletado com os dados atuais
			input.ID = existing.ID
			input.IsWatched = false
			input.PersonalRating = 0
			h.DB.Unscoped().Save(&input)
			h.DB.Unscoped().Model(&input).Update("deleted_at", nil)
			c.JSON(http.StatusCreated, input)
			return
		}
		// Registro ativo — já está na watchlist
		c.JSON(http.StatusConflict, gin.H{"error": "Filme já existe na watchlist"})
		return
	}

	if err := h.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar filme"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// UpdateMovie atualiza is_watched e/ou personal_rating de um filme
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		IsWatched      *bool `json:"is_watched"`
		PersonalRating *int  `json:"personal_rating"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var movie models.Movie
	if err := h.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
		return
	}

	updates := map[string]interface{}{}
	if input.IsWatched != nil {
		updates["is_watched"] = *input.IsWatched
	}
	if input.PersonalRating != nil {
		updates["personal_rating"] = *input.PersonalRating
	}

	h.DB.Model(&movie).Updates(updates)
	c.JSON(http.StatusOK, movie)
}

// DeleteMovie remove um filme da watchlist
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(&models.Movie{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir filme"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Filme excluído com sucesso"})
}

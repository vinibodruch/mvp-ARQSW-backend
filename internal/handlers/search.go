package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// omdbSearchResult representa o retorno da OMDb API
type omdbSearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Poster string `json:"Poster"`
	Response string `json:"Response"`
	Error  string `json:"Error"`
}

// SearchMovie busca um filme na OMDb API e retorna os dados formatados
func SearchMovie(omdbKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		if title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'title' é obrigatório"})
			return
		}

		url := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", url.QueryEscape(title), omdbKey)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao consultar OMDb"})
			return
		}
		defer resp.Body.Close()

		var result omdbSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao decodificar resposta"})
			return
		}

		if result.Response == "False" {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
		}

		// Retorna apenas os campos necessários para o front-end
		c.JSON(http.StatusOK, gin.H{
			"imdb_id":    result.ImdbID,
			"title":      result.Title,
			"year":       result.Year,
			"poster_url": result.Poster,
		})
	}
}

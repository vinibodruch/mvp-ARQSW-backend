package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ImdbID         string `gorm:"uniqueIndex;not null" json:"imdb_id"`
	Title          string `gorm:"not null"             json:"title"`
	Year           string `json:"year"`
	PosterURL      string `json:"poster_url"`
	IsWatched      bool   `gorm:"default:false"        json:"is_watched"`
	PersonalRating int    `gorm:"default:0"            json:"personal_rating"`
}

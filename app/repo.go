package app

import "github.com/shelik/vk-go-back/models"

// Repository ...
type Repository interface {
	Close() error
	GetGalleries(string) []models.Gallery
}

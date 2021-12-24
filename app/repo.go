package app

import "github.com/shelik/vk-go-back/models"

// Repository ...
type Repository interface {
	Close() error
	GetGalleries(ownerID, token string) []models.Gallery
	GetPhotos(ownerID, token string, galleryIDs []string, count int) []models.Photo
}

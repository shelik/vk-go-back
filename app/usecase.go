package app

import (
	"context"

	"github.com/shelik/vk-go-back/models"
)

// Usecase ...
type Usecase interface {
	GetGalleries(ctx context.Context, ownerID string) []models.Gallery
}

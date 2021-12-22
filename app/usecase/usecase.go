package usecase

import (
	"context"

	"github.com/shelik/vk-go-back/app"
	"github.com/shelik/vk-go-back/models"
)

// Usecase ...
type Usecase struct {
	repo app.Repository
}

// NewUsecase ...
func NewUsecase(repo app.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// GetGalleries ...
func (u *Usecase) GetGalleries(ctx context.Context, ownerID string) []models.Gallery {

	return u.repo.GetGalleries(ownerID)
}

// Translate ...
func (u *Usecase) GetPhotos(ctx context.Context, ownerID string, galleryIDs []string) []models.Photo {
	return u.repo.GetPhotos(ownerID, galleryIDs)
}

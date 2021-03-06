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
func (u *Usecase) GetGalleries(ctx context.Context, ownerID, token string) []models.Gallery {

	return u.repo.GetGalleries(ownerID, token)
}

// Translate ...
func (u *Usecase) GetPhotos(ctx context.Context, ownerID, token string, galleryIDs []string, count int) []models.Photo {
	return u.repo.GetPhotos(ownerID, token, galleryIDs, count)
}

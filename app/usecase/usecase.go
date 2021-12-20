package usecase

import (
	"context"

	"github.com/shelik/vk-go-back/app"
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

// Translate ...
func (u *Usecase) GetGalleries(ctx context.Context, ownerID string) []string {
	println("UC GetGalleries")
	return u.repo.GetGalleries(ownerID)
}

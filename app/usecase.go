package app

import "context"

// Usecase ...
type Usecase interface {
	GetGalleries(ctx context.Context, ownerID string) []string
}

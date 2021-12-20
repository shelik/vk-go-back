package app

// Repository ...
type Repository interface {
	Close() error
	GetGalleries(string) []string
}

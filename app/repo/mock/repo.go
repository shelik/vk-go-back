package mock

// RepoMock ...
type RepoMock struct {
}

// NewRepo ...
func NewRepo() *RepoMock {
	return &RepoMock{}
}

// Close ...
func (r *RepoMock) Close() error {
	return nil
}

// Translate ...
func (r *RepoMock) Translate(srcLang string, desLang string, text string) string {
	return "Mock Ok"
}

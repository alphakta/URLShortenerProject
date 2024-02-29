// /domain/url/service.go
package url

import (
	"github.com/alphakta/URLShortenerProject/repository/url"
)

type Service struct {
	Repo url.Repository
}

func NewService(repo url.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Create(longURL string) (string, error) {
	return s.Repo.AddURL(longURL)
}

func (s *Service) FindByShortURL(shortURL string) (string, error) {
	return s.Repo.FindLongURL(shortURL)
}

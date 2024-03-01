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

func (s *Service) GetTotalLinks() (int, error) {
	return s.Repo.GetTotalLinks()
}

func (s *Service) GetClicksByID(id string) (int, error) {
	return s.Repo.GetClicksByID(id)
}

func (s *Service) IncrementClicks(shortURL string) error {
	return s.Repo.IncrementClicks(shortURL)
}

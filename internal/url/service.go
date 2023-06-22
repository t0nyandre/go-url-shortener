package url

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/t0nyandre/go-url-shortener/internal/helpers"
	"github.com/teris-io/shortid"
)

type Service interface {
	GetLongUrl(shortUrl string) (*Url, error)
	Shorten(longUrl string) (*Url, error)
	IncrementClicks(url *Url) (*Url, error)
}

type service struct {
	repo   Repository
	logger *zerolog.Logger
}

// GetLongUrl implements Service.
func (s *service) GetLongUrl(shortUrl string) (*Url, error) {
	url, err := s.repo.GetByShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (s *service) IncrementClicks(url *Url) (*Url, error) {
	url.Clicks++
	url, err := s.repo.Update(url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

// Shorten implements Service.
func (s *service) Shorten(incoming string) (*Url, error) {
	longUrl := helpers.RemoveWhitespaces(incoming)
	urlStruct := &Url{LongUrl: longUrl}
	if urlStruct.LongUrl == "" {
		return nil, fmt.Errorf("you cannot provide an empty url")
	}

	shortid, err := shortid.Generate()
	if err != nil {
		return nil, err
	}
	urlStruct.ShortUrl = shortid

	url, err := s.repo.Create(urlStruct)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func NewService(repo Repository, logger *zerolog.Logger) Service {
	return &service{repo: repo, logger: logger}
}

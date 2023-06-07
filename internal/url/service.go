package url

import (
	"github.com/rs/zerolog"
	"github.com/teris-io/shortid"
)

type Service interface {
	Redirect(shortUrl string) (*Url, error)
	Shorten(longUrl string) (*Url, error)
}

type service struct {
	repo   Repository
	logger *zerolog.Logger
}

// Redirect implements Service.
func (s *service) Redirect(shortUrl string) (*Url, error) {
	url, err := s.repo.GetByShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}
	return url, nil
}

// Shorten implements Service.
func (s *service) Shorten(longUrl string) (*Url, error) {
	urlStruct := &Url{LongUrl: longUrl}
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

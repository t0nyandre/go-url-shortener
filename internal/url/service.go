package url

import (
	"github.com/rs/zerolog"
)

type Service interface {
	Redirect(incoming *Url) (*Url, error)
	Shorten(longUrl string) (*Url, error)
}

type service struct {
	repo   Repository
	logger *zerolog.Logger
}

// Redirect implements Service.
func (*service) Redirect(incoming *Url) (*Url, error) {
	panic("unimplemented")
}

// Shorten implements Service.
func (*service) Shorten(longUrl string) (*Url, error) {
	panic("unimplemented")
}

func NewService(repo Repository, logger *zerolog.Logger) Service {
	return &service{repo: repo, logger: logger}
}

package url

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Repository interface {
	GetByShortUrl(shortUrl string) (*Url, error)
	Create(url *Url) (*Url, error)
}

type repository struct {
	db     *sqlx.DB
	logger *zerolog.Logger
}

// Create implements Repository.
func (*repository) Create(url *Url) (*Url, error) {
	panic("unimplemented")
}

// GetByShortUrl implements Repository.
func (*repository) GetByShortUrl(shortUrl string) (*Url, error) {
	panic("unimplemented")
}

func NewRepository(db *sqlx.DB, logger *zerolog.Logger) Repository {
	return &repository{db: db, logger: logger}
}

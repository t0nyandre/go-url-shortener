package url

import (
	"errors"

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
func (repo *repository) Create(url *Url) (*Url, error) {
	rows, err := repo.db.Queryx(`INSERT INTO urls (short_url, long_url, clicks) VALUES ($1, $2, $3) RETURNING id`, url.ShortUrl, url.LongUrl, url.Clicks)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		url.ID = id
	}

	return url, nil
}

// GetByShortUrl implements Repository.
func (repo *repository) GetByShortUrl(shortUrl string) (*Url, error) {
	url := &Url{
		ShortUrl: shortUrl,
	}
	rows, err := repo.db.NamedQuery(`SELECT * FROM urls WHERE short_url = :short_url`, url)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.StructScan(url)
		if err != nil {
			return nil, err
		}
	}

	if url.ID == 0 {
		return nil, errors.New("url not found")
	}
	// TODO: Implement click counter
	return url, nil
}

func NewRepository(db *sqlx.DB, logger *zerolog.Logger) Repository {
	return &repository{db: db, logger: logger}
}

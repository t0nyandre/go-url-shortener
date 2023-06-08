package url

import (
	"errors"
	"fmt"

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
	result, err := repo.db.NamedExec(`INSERT INTO urls (short_url, long_url, clicks) VALUES (:short_url, :long_url, :clicks)`, url)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	url.ID = int32(lastId)

	return url, nil
}

// GetByShortUrl implements Repository.
func (repo *repository) GetByShortUrl(shortUrl string) (*Url, error) {
	url := &Url{
		ShortUrl: shortUrl,
	}
	test := &Url{}
	rows, err := repo.db.NamedQuery(`SELECT * FROM urls WHERE short_url = :short_url`, url)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("url not found")
	}

	fmt.Println(rows)
	for rows.Next() {
		err = rows.StructScan(&test)
		if err != nil {
			return nil, err
		}
	}
	// TODO: Implement click counter
	return test, nil
}

func NewRepository(db *sqlx.DB, logger *zerolog.Logger) Repository {
	return &repository{db: db, logger: logger}
}

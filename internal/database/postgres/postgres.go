package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/t0nyandre/go-url-shortener/internal/config"
)

func New(logger *zerolog.Logger, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		logger.Debug().Msgf("Failed to connect to database, will try again in 5 seconds: %v", err)
		time.Sleep(time.Duration(5) * time.Second)
		return New(logger, cfg)
	}

	logger.Info().Msg("Connected to database")
	return db, nil
}

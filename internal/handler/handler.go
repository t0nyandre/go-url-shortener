package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/t0nyandre/go-url-shortener/internal/config"
	"github.com/t0nyandre/go-url-shortener/internal/url"
)

type resource struct {
	logger *zerolog.Logger
	db     *sqlx.DB
	cfg    *config.Config
}

func NewHandler(logger *zerolog.Logger, db *sqlx.DB, cfg *config.Config) *resource {
	return &resource{
		logger: logger,
		db:     db,
		cfg:    cfg,
	}
}

func (res *resource) SetupHandlers() chi.Router {
	r := chi.NewRouter()
	// URL
	r.Mount("/api/v1/url", url.RegisterHandlers(res.db, res.logger))
	// Healthcheck
	r.Get("/api/_hc", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s: OK", res.cfg.Name)))
	})
	return r
}
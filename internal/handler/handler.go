package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
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
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// URL
	urlHandlers := url.RegisterHandlers(res.db, res.logger)
	r.Route("/v1/urls", func(r chi.Router) {
		r.Get("/{url}", urlHandlers.GetLongUrl)
		r.Post("/", urlHandlers.Shorten)
	})

	r.Route("/r", func(r chi.Router) {
		r.Get("/", urlHandlers.Redirect)
		r.Get("/{url}", urlHandlers.Redirect)
	})

	// Healthcheck
	r.Get("/_hc", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s: OK", res.cfg.Name)))
	})
	return r
}

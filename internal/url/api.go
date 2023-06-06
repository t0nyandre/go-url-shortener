package url

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func RegisterHandlers(db *sqlx.DB, logger *zerolog.Logger) chi.Router {
	repository := NewRepository(db, logger)
	service := NewService(repository, logger)

	res := resource{db: db, logger: logger, service: service}

	r := chi.NewRouter()
	// TODO: Create a way to delete short URLs
	r.Route("/{url}", func(r chi.Router) {
		r.Get("/", res.redirect)
	})
	r.Post("/", res.shorten)
	return r
}

type resource struct {
	db      *sqlx.DB
	logger  *zerolog.Logger
	service Service
}

func (res *resource) redirect(w http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

func (res *resource) shorten(w http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

package url

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
		r.Get("/", res.getLongUrl)
	})
	r.Post("/", res.shorten)
	return r
}

type resource struct {
	db      *sqlx.DB
	logger  *zerolog.Logger
	service Service
}

// TODO: Use render.Render
func (res *resource) Redirect(w http.ResponseWriter, req *http.Request) {
	shortUrl := chi.URLParam(req, "url")
	url, err := res.service.GetLongUrl(shortUrl)
	if err != nil {
		render.Render(w, req, &Response{
			HTTPStatusCode: http.StatusNotFound,
			Status:         "Failed",
			Errors: []ErrResponse{
				{
					Code:         "NOT_FOUND",
					ErrorMessage: err.Error(),
				},
			},
		})
		return
	}
	w.WriteHeader(http.StatusMovedPermanently)
	w.Header().Set("Location", url.LongUrl)
}

func (res *resource) getLongUrl(w http.ResponseWriter, req *http.Request) {
	shortUrl := chi.URLParam(req, "url")
	url, err := res.service.GetLongUrl(shortUrl)
	if err != nil {
		render.Render(w, req, &Response{
			HTTPStatusCode: http.StatusNotFound,
			Status:         "Failed",
			Errors: []ErrResponse{
				{
					Code:         "NOT_FOUND",
					ErrorMessage: err.Error(),
				},
			},
		})
		return
	}
	render.Render(w, req, &Response{
		HTTPStatusCode: http.StatusOK,
		Status:         "Success",
		Data:           url,
	})
}

func (res *resource) shorten(w http.ResponseWriter, req *http.Request) {
	urlStruct := &Url{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(urlStruct); err != nil {
		render.Render(w, req, &Response{
			HTTPStatusCode: http.StatusBadRequest,
			Status:         "Failed",
			Errors: []ErrResponse{
				{
					Code:         "BAD_REQUEST",
					ErrorMessage: err.Error(),
				},
			},
		})
		return
	}

	url, err := res.service.Shorten(urlStruct.LongUrl)
	if err != nil {
		render.Render(w, req, &Response{
			HTTPStatusCode: http.StatusInternalServerError,
			Status:         "Failed",
			Errors: []ErrResponse{
				{
					Code:         "INTERNAL_SERVER_ERROR",
					ErrorMessage: err.Error(),
				},
			},
		})
		return
	}

	render.Render(w, req, &Response{
		HTTPStatusCode: http.StatusCreated,
		Status:         "Success",
		Errors:         []ErrResponse{},
		Data:           url,
	})
}

type ErrResponse struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"message,omitempty"`
}

type Response struct {
	HTTPStatusCode int           `json:"-"`
	Status         string        `json:"status,omitempty"`
	Errors         []ErrResponse `json:"errors,omitempty"`
	Data           interface{}   `json:"data,omitempty"`
}

func (e *Response) Render(w http.ResponseWriter, req *http.Request) error {
	render.Status(req, e.HTTPStatusCode)
	return nil
}

package url

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/t0nyandre/go-url-shortener/internal/common"
)

func RegisterHandlers(db *sqlx.DB, logger *zerolog.Logger) *Resources {
	repository := NewRepository(db, logger)
	service := NewService(repository, logger)

	return &Resources{db: db, logger: logger, service: service}
}

type Resources struct {
	db      *sqlx.DB
	logger  *zerolog.Logger
	service Service
}

func (res *Resources) Redirect(w http.ResponseWriter, req *http.Request) {
	shortUrl := chi.URLParam(req, "url")
	if shortUrl == "" {
		render.Render(w, req, common.NewResponse(
			"Failed",
			http.StatusBadRequest,
			[]common.ErrResponse{
				{
					Code:         "BAD_REQUEST",
					ErrorMessage: "URL ID is required",
				},
			},
			nil,
		))
		return
	}
	url, err := res.service.GetLongUrl(shortUrl)
	if err != nil {
		render.Render(w, req, common.NewResponse(
			"Failed",
			http.StatusNotFound,
			[]common.ErrResponse{
				{
					Code:         "NOT_FOUND",
					ErrorMessage: err.Error(),
				},
			},
			nil,
		))
		return
	}
	url, err = res.service.IncrementClicks(url)
	if err != nil {
		render.Render(w, req, common.NewResponse(
			"Failed",
			http.StatusInternalServerError,
			[]common.ErrResponse{
				{
					Code:         "INTERNAL_SERVER_ERROR",
					ErrorMessage: err.Error(),
				},
			},
			nil,
		))
		return
	}
	http.Redirect(w, req, url.LongUrl, http.StatusMovedPermanently)
}

func (res *Resources) GetLongUrl(w http.ResponseWriter, req *http.Request) {
	shortUrl := chi.URLParam(req, "url")
	url, err := res.service.GetLongUrl(shortUrl)
	if err != nil {
		render.Render(w, req, common.NewResponse(
			"Failed",
			http.StatusNotFound,
			[]common.ErrResponse{
				{
					Code:         "NOT_FOUND",
					ErrorMessage: err.Error(),
				},
			},
			nil,
		))
		return
	}
	render.Render(w, req, common.NewResponse(
		"Success",
		http.StatusOK,
		[]common.ErrResponse{},
		url,
	))
}

func (res *Resources) Shorten(w http.ResponseWriter, req *http.Request) {
	urlStruct := &Url{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(urlStruct); err != nil {
		render.Render(w, req, common.NewResponse(
			"Failed",
			http.StatusBadRequest,
			[]common.ErrResponse{
				{
					Code:         "BAD_REQUEST",
					ErrorMessage: err.Error(),
				},
			},
			nil,
		))
		return
	}

	url, err := res.service.Shorten(urlStruct.LongUrl)
	if err != nil {
		render.Render(w, req, common.NewResponse(
			"Failed",
			http.StatusInternalServerError,
			[]common.ErrResponse{
				{
					Code:         "INTERNAL_SERVER_ERROR",
					ErrorMessage: err.Error(),
				},
			},
			nil,
		))
		return
	}

	render.Render(w, req, common.NewResponse(
		"Success",
		http.StatusCreated,
		[]common.ErrResponse{},
		url,
	))
}

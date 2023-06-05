package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/go-url-shortener/internal/config"
	"github.com/t0nyandre/go-url-shortener/internal/database/postgres"
	"github.com/t0nyandre/go-url-shortener/internal/logger"
)

func main() {
	var appConfig string

	flag.StringVar(&appConfig, "config", "./config/local.json", "path to config file")
	flag.Parse()

	l := logger.New()
	cfg, err := config.Load(appConfig)
	if err != nil {
		l.Debug().Msgf("Error loading config file. Using defaults: %v", err)
		cfg = config.Default()
	}

	_, err = postgres.New(l, cfg)
	if err != nil {
		l.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello world"))
	})

	l.Info().Msgf("Starting server on port %d\n", cfg.AppPort)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort), router); err != nil {
		l.Fatal().Msgf("Failed to start server: %s", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/t0nyandre/go-url-shortener/internal/config"
	"github.com/t0nyandre/go-url-shortener/internal/database/postgres"
	"github.com/t0nyandre/go-url-shortener/internal/handler"
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

	db, err := postgres.New(l, cfg)
	if err != nil {
		l.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	router := handler.NewHandler(l, db, cfg).SetupHandlers()

	l.Info().Msgf("Starting server on %s:%d", cfg.Host, cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), router); err != nil {
		l.Fatal().Msgf("Failed to start server: %s", err)
	}
}

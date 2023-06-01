package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/go-url-shortener/internal/config"
)

var appConfig = flag.String("config", "./config/local.json", "path to config file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*appConfig)
	if err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello world"))
	})

	fmt.Printf("Starting server on port %d\n", cfg.AppPort)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort), router); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

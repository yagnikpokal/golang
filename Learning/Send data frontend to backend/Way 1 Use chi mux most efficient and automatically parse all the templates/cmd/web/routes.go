package main

import (
	"net/http"
	"templatecache/pkg/config"
	"templatecache/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/service", handlers.Repo.Service)
	return r
}

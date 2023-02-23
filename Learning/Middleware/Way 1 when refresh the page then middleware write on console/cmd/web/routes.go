package main

import (
	"net/http"
	"templatecache/pkg/config"
	"templatecache/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(Writetoconsole)
	r.Use(NoSurf)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/service", handlers.Repo.Service)
	return r
}

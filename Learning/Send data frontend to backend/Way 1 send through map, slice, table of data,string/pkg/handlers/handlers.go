package handlers

import (
	"net/http"
	"templatecache/pkg/config"
	"templatecache/pkg/models"
	"templatecache/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"
	stringMap["state"] = "Gujarat"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Service is the handler for the home page
func (m *Repository) Service(w http.ResponseWriter, r *http.Request) {
	stringslice := []string{"This came from slice", "Yagnik", "Golang developer", "David", "JS developer"}
	render.RenderTemplate(w, "service.page.tmpl", &models.TemplateData{
		StringSlice: stringslice,
	})
}

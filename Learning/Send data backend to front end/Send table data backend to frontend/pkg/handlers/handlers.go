package handlers

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/models"
	"bookings-udemy/pkg/render"
	"net/http"
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
	stringMap := make(map[string]string)

	stringMap["Company_1"] = "Alfreds Futterkiste"
	stringMap["Contact_1"] = "Maria Anders"
	stringMap["Country_1"] = "Germany"

	stringMap["Company_2"] = "Centro comercial Moctezuma"
	stringMap["Contact_2"] = "Francisco Chang"
	stringMap["Country_2"] = "Mexico"

	stringMap["Company_3"] = "Ernst Handel"
	stringMap["Contact_3"] = "Maria Anders"
	stringMap["Country_3"] = "Austria"

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", nil)
}

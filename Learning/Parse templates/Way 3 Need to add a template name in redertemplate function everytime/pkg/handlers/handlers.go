package handlers

import (
	"net/http"
	"yagniktemplaterender/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// About is the handler for the about page
func Service(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "service.page.tmpl")
}

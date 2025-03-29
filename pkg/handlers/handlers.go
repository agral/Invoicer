package handlers

import (
	"Invoicer/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.tmpl")
}

func Status(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "status.tmpl")
}

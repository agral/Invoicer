package handlers

import (
	"Invoicer/pkg/config"
	"Invoicer/pkg/models"
	"Invoicer/pkg/render"
	"net/http"
)

// Repository pattern
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for individual page handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) Status(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Test"

	// send the data to the template
	render.RenderTemplate(w, "status.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

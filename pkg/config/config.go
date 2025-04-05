package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	IsProduction   bool
	UseCache       bool
	SessionManager *scs.SessionManager
	TemplateCache  map[string]*template.Template
}

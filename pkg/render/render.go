package render

import (
	"Invoicer/pkg/config"
	"Invoicer/pkg/models"
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var appConfig *config.AppConfig

func SetAppConfig(cfg *config.AppConfig) {
	appConfig = cfg
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var cache map[string]*template.Template
	var err error

	if appConfig.UseCache {
		// Use the template cache from the app config:
		cache = appConfig.TemplateCache
	} else {
		cache, err = CreateTemplateCache()
		if err != nil {
			log.Println(err)
		}
	}

	// Look up the requested template in the cache
	template, isOk := cache[tmpl]
	if !isOk {
		log.Printf("Template %s missing in cache", tmpl)
		log.Fatal("Aborting.")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err = template.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	// Parse all the files from `./templates/` ending with `.page.tmpl`:
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = templateSet
	}
	return cache, nil
}

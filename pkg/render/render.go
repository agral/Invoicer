package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate3(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %q: %s\n", tmpl, err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	var tmpl *template.Template
	var err error
	// if templates are in cache, use the cached versions:
	_, isCached := templateCache[templateName]
	if isCached {
		log.Printf("Using cached template %q", templateName)
	} else {
		log.Printf("Creating and caching template %q", templateName)
		err = createTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}
	}

	tmpl = templateCache[templateName]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %q: %s\n", templateName, err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[t] = tmpl
	return nil
}

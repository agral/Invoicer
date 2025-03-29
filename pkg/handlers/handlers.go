package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %q: %s\n", tmpl, err)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.tmpl")
}

func Status(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "status.tmpl")
}

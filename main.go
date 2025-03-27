package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/signintech/gopdf"
)

const BR_SIZE float64 = 13.5
const PORT_NUMBER string = ":31337"

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %q: %s\n", tmpl, err)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.tmpl")
}

func Status(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "status.tmpl")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/status", Status)

	fmt.Printf("Starting the web interface on http://localhost%s ...\n", PORT_NUMBER)
	_ = http.ListenAndServe(PORT_NUMBER, nil)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: gopdf.Rect{W: tp(210), H: tp(297)}, // A4, portrait
	})
	pdf.AddPage()
	if !loadFonts(&pdf) {
		fmt.Printf("Failed to load fonts. Exiting.")
		os.Exit(1)
	}
	err := pdf.SetFont("NotoSans", "", 10)
	b, err := os.ReadFile("data/default/header.json")
	if err != nil {
		fmt.Println("Failed to read data/default/header.json.")
		fmt.Printf("Error: %s\n", err)
	}
	var header Header
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&header)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	WriteRichText(&pdf, header.Left)
	WriteRichText(&pdf, header.Right)
	pdf.WritePdf("out.pdf")
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"Invoicer/pkg/config"
	"Invoicer/pkg/handlers"
	"Invoicer/pkg/render"

	"github.com/alexedwards/scs/v2"
	"github.com/signintech/gopdf"
)

const BR_SIZE float64 = 13.5
const PORT_NUMBER string = ":31337"

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {
	app.IsProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.IsProduction
	app.SessionManager = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create the template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false // Development mode: set to false; pages won't be cached.
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.SetAppConfig(&app)

	fmt.Printf("Starting the web interface on http://localhost%s ...\n", PORT_NUMBER)
	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: gopdf.Rect{W: tp(210), H: tp(297)}, // A4, portrait
	})
	pdf.AddPage()
	if !loadFonts(&pdf) {
		fmt.Printf("Failed to load fonts. Exiting.")
		os.Exit(1)
	}
	err = pdf.SetFont("NotoSans", "", 10)
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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/signintech/gopdf"
)

const BR_SIZE float64 = 13.5
const PORT_NUMBER string = ":31337"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Invoicer's web interface")
	if err != nil {
		fmt.Println(err)
	}
}
func Status(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Status page")
	if err != nil {
		fmt.Println(err)
	}
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

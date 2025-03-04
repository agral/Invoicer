package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/signintech/gopdf"
)

const BR_SIZE float64 = 13.5

type Header struct {
	Left []string `json:"left"`
}

func main() {
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
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pdf.SetY(tp(20))
	writeMultiLineText(&pdf, []string{"line1", "line2", "line3", "line4", "line5"}, tp(20), BR_SIZE)
	b, err := os.ReadFile("data/default/header.json")
	if err != nil {
		fmt.Println("Failed to read data/default/header.json.")
		fmt.Printf("Error: %s\n", err)
	}
	var header Header
	json.NewDecoder(bytes.NewBuffer(b)).Decode(&header)
	for _, h := range header.Left {
		fmt.Println(h)
	}
	fmt.Println(header)
	pdf.WritePdf("out.pdf")
}

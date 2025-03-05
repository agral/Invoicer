package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/signintech/gopdf"
)

const BR_SIZE float64 = 13.5

func WriteRichText(pdf *gopdf.GoPdf, richtext RichText) bool {
	err := pdf.SetFont(richtext.Font.Name, "", richtext.Font.Size)
	if err != nil {
		fmt.Printf("Error while calling SetFont(name=%q, size=%f) in WriteRichText", richtext.Font.Name, richtext.Font.Size)
		fmt.Printf("Error: %s\n", err)
		return false
	}
	pdf.SetY(tp(richtext.Position.Y))
	writeMultiLineText(pdf, richtext.Text, tp(richtext.Position.X), richtext.Font.BrSize)
	return true
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

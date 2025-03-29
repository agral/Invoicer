package main

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func loadFonts(pdf *gopdf.GoPdf) bool {
	err := pdf.AddTTFFont("NotoSans", "assets/ttf/NotoSans-Regular.ttf")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func writeMultiLineText(pdf *gopdf.GoPdf, lines []string, x float64, br float64) {
	for _, line := range lines {
		pdf.SetX(x)
		pdf.Cell(nil, line)
		pdf.Br(br)
	}
}

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

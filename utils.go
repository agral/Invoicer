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

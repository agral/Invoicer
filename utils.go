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

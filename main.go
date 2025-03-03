package main

import (
	"fmt"
	"os"

	"github.com/signintech/gopdf"
)

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
	pdf.SetXY(tp(20), tp(20))
	pdf.Cell(nil, "PDF file generated by Invoicer (now with loading of fonts offloaded to utils!)")
	pdf.WritePdf("out.pdf")
}

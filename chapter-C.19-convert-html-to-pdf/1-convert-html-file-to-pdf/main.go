package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
)

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("./input.html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}

package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"net/http"
	"time"
)

func startWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		w.Write([]byte(`
            <!DOCTYPE html>
            <html>
                <head>
                    <title>Testing</title>
                </head>
            <body>
                <p>Hello world!</p>
            </body>
            </html>
        `))
	})
	http.ListenAndServe(":9000", nil)
}

func main() {
	go startWebServer()
	time.Sleep(5 * time.Second)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	page := wkhtmltopdf.NewPage("http://localhost:9000")
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	pdfg.AddPage(page)

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

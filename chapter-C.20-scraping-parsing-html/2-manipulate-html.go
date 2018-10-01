package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/yosssi/gohtml"
	"log"
	"strings"
)

const sampleHTML = `<!DOCTYPE html>
    <html>
        <head>
            <title>Sample HTML</title>
        </head>
        <body>
            <h1>Header</h1>
            <footer class="footer-1">Footer One</footer>
            <footer class="footer-2">Footer Two</footer>
            <footer class="footer-3">Footer Three</footer>
        </body>
    </html>
`

func main() {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTML))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("h1").AfterHtml("<p>Lorem Ipsum Dolor Sit Amet Gedhang Goreng</p>")
	doc.Find("p").AppendHtml(" <b>Tournesol</b>")
	doc.Find("h1").SetAttr("class", "header")
	doc.Find("footer").First().Remove()
	doc.Find("body > *:nth-child(4)").Remove()

	modifiedHTML, err := goquery.OuterHtml(doc.Selection)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(gohtml.Format(modifiedHTML))
}

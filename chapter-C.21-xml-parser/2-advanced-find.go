package main

import (
	"github.com/beevik/etree"
	"log"
)

type M map[string]interface{}

func main() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./data.xml"); err != nil {
		log.Fatal(err.Error())
	}

	root := doc.SelectElement("website")

	popularArticleText := root.FindElement(`//article[@info='popular article']/title`)
	if popularArticleText != nil {
		log.Println("Popular article", popularArticleText.Text())
	}
}

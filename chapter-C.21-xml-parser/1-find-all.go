package main

import (
	"encoding/json"
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
	rows := make([]M, 0)

	for _, article := range root.FindElements("//article") {
		row := make(M)
		row["title"] = article.SelectElement("title").Text()
		row["url"] = article.SelectElement("url").Text()

		categories := make([]string, 0)
		for _, category := range article.SelectElements("category") {
			categories = append(categories, category.Text())
		}
		row["categories"] = categories

		if info := article.SelectAttr("info"); info != nil {
			row["info"] = info.Value
		}

		rows = append(rows, row)
	}

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}

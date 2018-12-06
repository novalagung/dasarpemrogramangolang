package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	bookNameFlag := flag.String("name", "", "Book Name")
	flag.Parse()

	bookName := *bookNameFlag
	if bookName == "" {
		log.Fatal("-name argument is required")
		return
	}

	regex := regexp.MustCompile(`<title>(.*?)<\/title>`)

	basePath, _ := os.Getwd()
	bookPath := filepath.Join(basePath, "_book")

	err := filepath.Walk(bookPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) != ".html" {
			return nil
		}

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		htmlString := string(buf)

		title := regex.FindString(htmlString)
		title = strings.Replace(title, "· GitBook", fmt.Sprintf("- %s", bookName), -1)
		if strings.HasPrefix(title, "Golang") {
			// do nothing
		} else if strings.HasPrefix(title, "Go") {
			title = strings.Replace(title, "Go", "Golang", 1)
		} else {
			title = fmt.Sprintf("Golang %s", title)
		}
		newHtmlString := strings.Replace(htmlString, title, title, -1)

		err = ioutil.WriteFile(path, []byte(newHtmlString), info.Mode())
		if err != nil {
			return err
		}

		fmt.Println("  ==>", path)

		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	siteMapPath := filepath.Join(basePath, "_book", "sitemap.xml")

	buf, err := ioutil.ReadFile(siteMapPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	htmlString := string(buf)
	newHtmlString := strings.Replace(htmlString, `<changefreq>weekly</changefreq>`, `<changefreq>daily</changefreq>`, -1)

	err = ioutil.WriteFile(siteMapPath, []byte(newHtmlString), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("  ==>", siteMapPath)
}

//

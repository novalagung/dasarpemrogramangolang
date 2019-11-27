package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	bookName := "Dasar Pemrograman Golang"
	adClient := "ca-pub-1417781814120840"

	regex := regexp.MustCompile(`<title>(.*?)<\/title>`)

	basePath, _ := os.Getwd()
	bookPath := filepath.Join(basePath, "_book")

	err := filepath.Walk(bookPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
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

		oldTitle := regex.FindString(htmlString)
		oldTitle = strings.Replace(oldTitle, "<title>", "", -1)
		oldTitle = strings.Replace(oldTitle, "</title>", "", -1)

		isLandingPage := false
		newTitle := oldTitle
		if newTitle == "Introduction · GitBook" {
			isLandingPage = true
			newTitle = bookName
		} else {
			if titleParts := strings.Split(newTitle, "."); len(titleParts) > 2 {
				actualTitle := strings.TrimSpace(titleParts[2])

				if strings.Contains(actualTitle, "Go") || strings.Contains(actualTitle, "Golang") {
					// do nothing
				} else {
					titleParts[2] = fmt.Sprintf(" Golang %s", actualTitle)
				}

				newTitle = strings.Join(titleParts, ".")
			}

			newTitle = strings.Replace(newTitle, "· GitBook", fmt.Sprintf("- %s", bookName), -1)
		}
		htmlString = strings.Replace(htmlString, oldTitle, newTitle, -1)

		metaToFind := `<meta content=""name="description">`
		metaReplacement := metaToFind
		if isLandingPage {
			metaReplacement = `<meta content="Belajar Pemrograman Go Mulai Dari 0" name="description">`
		}

		htmlString = strings.Replace(htmlString, metaToFind, fmt.Sprintf(`%s<script data-ad-client="%s" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script><script>(adsbygoogle = window.adsbygoogle || []).push({ google_ad_client: "%s", enable_page_level_ads: true }); </script>`, metaReplacement, adClient, adClient), -1)

		err = ioutil.WriteFile(path, []byte(htmlString), info.Mode())
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

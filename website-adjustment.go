package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var BASE_VERSION = 2

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

		// ==== remove invalid lang tag for EPUB validation
		htmlString = strings.Replace(htmlString, ` lang="" xml:lang=""`, "", -1)

		// ==== remove colon from id for EPUB validation
		// re, err := regexp.Compile(`id=\"(.*?)\"`)
		// if err != nil {
		// 	return err
		// }
		// for _, each := range re.FindAllString(htmlString, -1) {
		// 	fmt.Println(each)
		// 	newID := strings.Replace(each, "-", "_", -1)
		// 	htmlString = strings.Replace(htmlString, each, newID, -1)
		// }

		// ==== adjust title for SEO purpose
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

		// ==== adjust meta for SEO purpose
		metaToFind := `<meta content=""name="description">`
		metaReplacement := metaToFind
		if isLandingPage {
			metaReplacement = `<meta content="Belajar Pemrograman Go Mulai Dari 0" name="description">`
		}
		htmlString = strings.Replace(htmlString, metaToFind, fmt.Sprintf(`%s<script data-ad-client="%s" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script><script>(adsbygoogle = window.adsbygoogle || []).push({ google_ad_client: "%s", enable_page_level_ads: true }); </script>`, metaReplacement, adClient, adClient), -1)

		// ==== inject github stars button
		buttons := `<div style="position: fixed; top: 10px; right: 30px; padding: 10px; background-color: rgba(255, 255, 255, 0.7);">
			<a class="github-button" href="https://github.com/novalagung/dasarpemrogramangolang" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star novalagung/dasarpemrogramangolang on GitHub">Star</a>
			&nbsp;
			<a class="github-button" href="https://github.com/novalagung" data-size="large" aria-label="Follow @novalagung on GitHub">Follow @novalagung</a>
			<script async defer src="https://buttons.github.io/buttons.js"></script>
		</div>`
		htmlString = strings.Replace(htmlString, `</body>`, fmt.Sprintf("%s</body>", buttons), -1)

		newsSection := `<a href="https://devops.novalagung.com/en/" target="_blank" class="book-news">Hi all, I just released another ebook. This new one focus on devops implementation of various stacks rather than specific programming topic. Only few articles have been published, more are coming. Have a look! https://devops.novalagung.com/en/</a>`
		htmlString = strings.Replace(htmlString, `</body>`, fmt.Sprintf(`%s</body>`, newsSection), -1)

		buttonScript := `<script async defer src="https://buttons.github.io/buttons.js"></script>`
		htmlString = strings.Replace(htmlString, `</head>`, fmt.Sprintf("%s</head>", buttonScript), -1)

		// ==== update file
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

	// index adjustment
	indexPath := filepath.Join(basePath, "_book", "index.html")

	buf, err := ioutil.ReadFile(indexPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	indexContent := strings.Replace(string(buf), `((VERSION))`, fmt.Sprintf("%d.%s", BASE_VERSION, time.Now().Format("2006.01.02.150405")), -1)

	// update file
	err = ioutil.WriteFile(indexPath, []byte(indexContent), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	// sitemap adjustment
	siteMapPath := filepath.Join(basePath, "_book", "sitemap.xml")

	buf, err = ioutil.ReadFile(siteMapPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	sitemapContent := strings.Replace(string(buf), `<changefreq>weekly</changefreq>`, `<changefreq>daily</changefreq>`, -1)

	err = ioutil.WriteFile(siteMapPath, []byte(sitemapContent), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("  ==>", siteMapPath)
}

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
	"time"
)

var (
	baseVersion = 2
	bookName    = "Dasar Pemrograman Golang"
	adClient    = "ca-pub-1417781814120840"
)

func main() {
	flagAdjustment := flag.String("type", "", "adjustment type (pre/post)")
	flag.Parse()

	switch *flagAdjustment {
	case "pre":
		preAdjustment()
	case "post":
		postAdjustment()
	default:
		break
	}
}

func preAdjustment() {
	basePath, _ := os.Getwd()
	readmePath := filepath.Join(basePath, "README.md")

	buf, err := ioutil.ReadFile(readmePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	mdString := string(buf)

	// ==== adjust version
	versionToFind := `((VERSION))`
	versionReplacement := fmt.Sprintf("%d.%s", baseVersion, time.Now().Format("2006.01.02.150405"))
	mdString = strings.Replace(mdString, versionToFind, versionReplacement, -1)

	err = ioutil.WriteFile(readmePath, []byte(mdString), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func postAdjustment() {
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

		// ==== adjust title for SEO purpose
		oldTitle := regex.FindString(htmlString)
		oldTitle = strings.Replace(oldTitle, "<title>", "", -1)
		oldTitle = strings.Replace(oldTitle, "</title>", "", -1)
		newTitle := oldTitle
		isLandingPage := oldTitle == "Introduction · GitBook"
		if isLandingPage {
			newTitle = bookName
		} else {
			if titleParts := strings.Split(newTitle, "."); len(titleParts) > 2 {
				actualTitle := strings.TrimSpace(titleParts[2])

				if !(strings.Contains(actualTitle, "Go") || strings.Contains(actualTitle, "Golang")) {
					titleParts[2] = fmt.Sprintf(" Golang %s", actualTitle)
				}
				newTitle = strings.Join(titleParts, ".")
			}
			newTitle = strings.Replace(newTitle, "· GitBook", fmt.Sprintf("- %s", bookName), -1)
		}
		htmlString = strings.Replace(htmlString, oldTitle, newTitle, -1)

		// ==== adjust meta for SEO purpose
		metaToFind := `<meta content=""name="description">`
		metaReplacement := ""
		if isLandingPage {
			metaReplacement = `<meta content="Belajar Pemrograman Go Mulai Dari 0" name="description">`
		}
		metaReplacement = metaReplacement + `<script data-ad-client="` + adClient + `" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script><script>(adsbygoogle = window.adsbygoogle || []).push({ google_ad_client: "` + adClient + `", enable_page_level_ads: true }); </script>`
		htmlString = strings.Replace(htmlString, metaToFind, metaReplacement, -1)

		// ==== inject github stars button
		buttonToFind := `</body>`
		buttonReplacement := `<div style="position: fixed; top: 10px; right: 30px; padding: 10px; background-color: rgba(255, 255, 255, 0.7);"><a class="github-button" href="https://github.com/novalagung/dasarpemrogramangolang" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star novalagung/dasarpemrogramangolang on GitHub">Star</a>&nbsp;<a class="github-button" href="https://github.com/novalagung" data-size="large" aria-label="Follow @novalagung on GitHub">Follow @novalagung</a><script async defer src="https://buttons.github.io/buttons.js"></script></div>`+buttonToFind
		htmlString = strings.Replace(htmlString, buttonToFind, buttonReplacement, -1)

		// ==== inject github stars js script
		buttonScriptToFind := `</head>`
		buttonScriptReplacement := `<script async defer src="https://buttons.github.io/buttons.js"></script>`+buttonScriptToFind
		htmlString = strings.Replace(htmlString, buttonScriptToFind, buttonScriptReplacement, -1)

		// ===== inject banner of new ebook
		bannerToFind := `</body>`
		bannerReplacement := `<a href="https://devops.novalagung.com/" target="_blank" class="book-news">Halo semua, Saya telah merilis ebook baru lo, tentang devops. Di ebook tersebut fokus tentang pembahasan banyak sekali stacks/teknologi devops, jadi tidak hanya membahas satu stack saja. Dan kabar baiknya tersedia dalam dua bahasa, Indonesia dan Inggris. Yuk mampir https://devops.novalagung.com/</a>`+bannerToFind
		htmlString = strings.Replace(htmlString, bannerToFind, bannerReplacement, -1)

		// ===== inject info banner if exists
		infoBannerToFind := `</body>`
		infoBannerReplacement := `<div class="banner-container" onclick="this.style.display = 'none';"><div><a target="_blank" href="https://www.udemy.com/course/praktis-belajar-docker-dan-kubernetes-untuk-pemula/"><img src="/images/banner.png"></a></div></div><script>var bannerCounter = parseInt(localStorage.getItem("banner-counter")); if (isNaN(bannerCounter)) { bannerCounter = 0; } var bannerEl = document.querySelector('.banner-container'); if (bannerCounter % 5 === 0 && bannerEl) { bannerEl.style.display = 'block'; } localStorage.setItem("banner-counter", String(bannerCounter + 1));</script>`+infoBannerToFind
		htmlString = strings.Replace(htmlString, infoBannerToFind, infoBannerReplacement, -1)

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

	// ===== sitemap adjustment
	siteMapPath := filepath.Join(basePath, "_book", "sitemap.xml")
	buf, err := ioutil.ReadFile(siteMapPath)
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

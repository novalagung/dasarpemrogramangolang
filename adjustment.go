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
	baseVersion = 3
	bookName    = "Dasar Pemrograman Golang"
	adClient    = "ca-pub-1417781814120840"
	now         = time.Now()
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
	mdString = strings.Replace(mdString, versionToFind, getVersion(), -1)

	err = ioutil.WriteFile(readmePath, []byte(mdString), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func postAdjustment() {
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
		oldTitle := getTitle(htmlString)
		newTitle := oldTitle
		isLandingPage := oldTitle == "Introduction · GitBook"
		if isLandingPage {
			newTitle = bookName
		} else {
			// ==== reformat title into "<title-name> - <ebook-name>"
			titleParts := strings.Split(newTitle, ".")
			if len(titleParts) > 2 {
				actualTitle := strings.TrimSpace(titleParts[2])

				if !(strings.Contains(actualTitle, "Go") || strings.Contains(actualTitle, "Golang")) {
					titleParts[2] = fmt.Sprintf(" Golang %s", actualTitle)
				}
				newTitle = strings.Join(titleParts, ".")
			}
			if newTitle == "Belajar Golang" {
				newTitle = "Tutorial Gratis Belajar Dasar Pemrograman Golang Mulai Dari 0"
			}
			newTitle = strings.Replace(newTitle, "· GitBook", fmt.Sprintf("- %s", bookName), -1)

			// ==== remove the "A.2"-ish from the title
			titleParts = strings.Split(newTitle, " ")
			if strings.Contains(titleParts[0], ".") {
				titleParts = titleParts[1:len(titleParts)]
			}
			newTitle = strings.Join(titleParts, " ")
		}
		htmlString = strings.Replace(htmlString, oldTitle, newTitle, -1)

		// ==== adjust meta for SEO purpose
		metaReplacement := ""
		if isLandingPage {
			metaReplacement = `<meta content="Tutorial Gratis Belajar Dasar Pemrograman Golang Mulai Dari 0" name="description">`
		}
		htmlString = strings.Replace(htmlString, `<meta name="description" content="">`, metaReplacement, -1)

		// ==== images' alt
		imagesAltToFind := ` alt="`
		imagesAltReplacement := imagesAltToFind + bookName + " - "
		htmlString = strings.Replace(htmlString, imagesAltToFind, imagesAltReplacement, -1)

		// ==== disqus lazy load
		disqusJSBuf, _ := ioutil.ReadFile("./gitbook-plugin-disqus.js")
		ioutil.WriteFile("./_book/gitbook/gitbook-plugin-disqus/plugin.js", disqusJSBuf, 0644)

		// ==== gitbook assets lazy load
		cssToLoad := []string{
			// "gitbook/style.css",
			"gitbook/gitbook-plugin-disqus/plugin.css",
			"gitbook/gitbook-plugin-highlight/website.css",
			"gitbook/gitbook-plugin-search/search.css",
			"gitbook/gitbook-plugin-fontsettings/website.css",
			// `/adjustment.css?v=` + getVersion() + `"`,
		}
		for _, cssFileNameToFind := range cssToLoad {
			cssFileNameReplacement := fmt.Sprintf(`%s" media="print" onload="this.media='all'`, cssFileNameToFind)
			htmlString = strings.Replace(htmlString, cssFileNameToFind, cssFileNameReplacement, -1)
		}

		// ==== inject github stars button
		buttonToFind := `</body>`
		buttonReplacement := `<div style="position: fixed; top: 10px; right: 30px; padding: 10px; background-color: rgba(255, 255, 255, 0.7);"><a class="github-button" href="https://github.com/novalagung/dasarpemrogramangolang" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star novalagung/dasarpemrogramangolang on GitHub">Star</a>&nbsp;<a class="github-button" href="https://github.com/novalagung" data-size="large" aria-label="Follow @novalagung on GitHub">Follow @novalagung</a><script async defer src="https://buttons.github.io/buttons.js"></script></div>` + buttonToFind
		htmlString = strings.Replace(htmlString, buttonToFind, buttonReplacement, -1)

		// ==== inject adjustment css
		adjustmentCSSBuf, _ := ioutil.ReadFile("./adjustment.css")
		ioutil.WriteFile("./_book/gitbook/adjustment.css", adjustmentCSSBuf, 0644)
		adjustmentCSSToFind := `</head>`
		adjustmentCSSReplacement := `<link rel="stylesheet" href="gitbook/adjustment.css?v=` + getVersion() + `">` + adjustmentCSSToFind
		htmlString = strings.Replace(htmlString, adjustmentCSSToFind, adjustmentCSSReplacement, -1)

		// ==== inject github stars js script
		buttonScriptToFind := `</head>`
		buttonScriptReplacement := `<script async defer src="https://buttons.github.io/buttons.js"></script>` + buttonScriptToFind
		htmlString = strings.Replace(htmlString, buttonScriptToFind, buttonScriptReplacement, -1)

		// ==== google ads
		// googleAdsToFind := `</head>`
		// // googleAdsReplacement := `<script data-ad-client="` + adClient + `" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js">` + `</script><script>(adsbygoogle = window.adsbygoogle || []).push({ google_ad_client: "` + adClient + `", enable_page_level_ads: true }); </script>` + googleAdsToFind
		// googleAdsReplacement := `<script data-ad-client="` + adClient + `" async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>` + googleAdsToFind
		// htmlString = strings.Replace(htmlString, googleAdsToFind, googleAdsReplacement, -1)

		// ===== inject fb pixel
		// fbPixelToFind := `</head>`
		// fbPixelReplacement := `<script>!function(f,b,e,v,n,t,s){if(f.fbq)return;n=f.fbq=function(){n.callMethod?n.callMethod.apply(n,arguments):n.queue.push(arguments)};if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';n.queue=[];t=b.createElement(e);t.async=!0;t.src=v;s=b.getElementsByTagName(e)[0];s.parentNode.insertBefore(t,s)}(window,document,'script','https://connect.facebook.net/en_US/fbevents.js');fbq('init','1247398778924723');fbq('track','PageView');</script><noscript><imgheight="1"width="1"style="display:none"src="https://www.facebook.com/tr?id=1247398778924723&ev=PageView&noscript=1"/></noscript>` + fbPixelToFind
		// htmlString = strings.Replace(htmlString, fbPixelToFind, fbPixelReplacement, -1)

		// ===== inject banner of new ebook
		// bannerToFind := `</body>`
		// bannerReplacement := `<a href="https://devops.novalagung.com/" target="_blank" class="book-news">Halo semua, Saya telah merilis ebook baru lo, tentang devops. Di ebook tersebut fokus tentang pembahasan banyak sekali stacks/teknologi devops, jadi tidak hanya membahas satu stack saja. Dan kabar baiknya tersedia dalam dua bahasa, Indonesia dan Inggris. Yuk mampir https://devops.novalagung.com/</a>` + bannerToFind
		// htmlString = strings.Replace(htmlString, bannerToFind, bannerReplacement, -1)

		// ===== inject popup info banner if exists
		// infoBannerToFind := `</body>`
		// infoBannerReplacement := `<div class="banner-container" onclick="this.style.display = 'none';"><div><a target="_blank" href="https://www.udemy.com/course/praktis-belajar-docker-dan-kubernetes-untuk-pemula/"><img src="/images/banner.png?v=` + getVersion() + `"></a></div></div><script>var bannerCounter = parseInt(localStorage.getItem("banner-counter")); if (isNaN(bannerCounter)) { bannerCounter = 0; } var bannerEl = document.querySelector('.banner-container'); if (bannerCounter % 5 === 1 && bannerEl) { bannerEl.style.display = 'block'; } localStorage.setItem("banner-counter", String(bannerCounter + 1));</script>` + infoBannerToFind
		// htmlString = strings.Replace(htmlString, infoBannerToFind, infoBannerReplacement, -1)

		// ==== update file
		err = ioutil.WriteFile(path, []byte(strings.TrimSpace(htmlString)), info.Mode())
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

	// ===== change crawl frequency
	sitemapContent := strings.Replace(string(buf), `<changefreq>weekly</changefreq>`, `<changefreq>daily</changefreq>`, -1)

	// ===== inject files into sitemap
	sitemapContent = strings.Replace(`</urlset>`, strings.TrimSpace(`
	<url>
		<loc>https://dasarpemrogramangolang.novalagung.com/dasarpemrogramangolang.pdf</loc>
		<changefreq>daily</changefreq>
		<priority>0.5</priority>
	</url>
	<url>
		<loc>https://dasarpemrogramangolang.novalagung.com/dasarpemrogramangolang.epub</loc>
		<changefreq>daily</changefreq>
		<priority>0.5</priority>
	</url>
	<url>
		<loc>https://dasarpemrogramangolang.novalagung.com/dasarpemrogramangolang.mobi</loc>
		<changefreq>daily</changefreq>
		<priority>0.5</priority>
	</url>
</urlset>`))

	err = ioutil.WriteFile(siteMapPath, []byte(sitemapContent), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("  ==>", siteMapPath)
}

func getVersion() string {
	return fmt.Sprintf("%d.%s", baseVersion, now.Format("2006.01.02.150405"))
}

func getTitle(htmlString string) string {
	regexFindTitle := regexp.MustCompile(`<title>(.*?)<\/title>`)
	title := regexFindTitle.FindString(htmlString)
	title = strings.Replace(title, "<title>", "", -1)
	title = strings.Replace(title, "</title>", "", -1)

	return title
}

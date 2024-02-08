package main

import (
	"flag"
	"fmt"
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
	ga4tagId    = "G-MZ74P74K72"
	now         = time.Now()
)

func main() {
	flagMode := flag.String("mode", "", "mode (webbook/ebook)")
	flagAdjustment := flag.String("adjustment", "", "adjustment type (pre/post)")
	flag.Parse()

	switch *flagMode {
	case "webbook":
		switch *flagAdjustment {
		case "pre":
			webbookPreAdjustment()
		case "post":
			webbookPostAdjustment()
		default:
			log.Fatalf("unrecognized flag -adjustment")
		}
	case "ebook":
		switch *flagAdjustment {
		case "pre":
			ebookPreAdjustment()
		default:
			log.Fatalf("unrecognized flag -adjustment")
		}
	default:
		log.Fatalf("unrecognized flag -mode")
	}
}

func getVersion() string {
	return fmt.Sprintf("%d.%s", baseVersion, now.Format("2006.01.02.150405"))
}

func webbookPreAdjustment() {
	basePath, _ := os.Getwd()
	readmePath := filepath.Join(basePath, "README.md")

	buf, err := os.ReadFile(readmePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	mdString := string(buf)

	// ==== adjust version
	versionToFind := `((VERSION))`
	mdString = strings.ReplaceAll(mdString, versionToFind, getVersion())

	err = os.WriteFile(readmePath, []byte(mdString), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func webbookPostAdjustment() {
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

		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		htmlString := string(buf)

		// ==== remove invalid lang tag for EPUB validation
		htmlString = strings.ReplaceAll(htmlString, ` lang="" xml:lang=""`, "")

		// ==== adjust title for SEO purpose
		oldTitle := func(htmlString string) string {
			regexFindTitle := regexp.MustCompile(`<title>(.*?)<\/title>`)
			title := regexFindTitle.FindString(htmlString)
			title = strings.ReplaceAll(title, "<title>", "")
			title = strings.ReplaceAll(title, "</title>", "")
			return title
		}(htmlString)

		newTitle := oldTitle
		isLandingPage := (oldTitle == "Introduction · GitBook") || (oldTitle == "Introduction · HonKit") || (oldTitle == "Introduction &#xB7; GitBook") || (oldTitle == "Introduction &#xB7; HonKit")
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
			newTitle = strings.ReplaceAll(newTitle, "· GitBook", fmt.Sprintf("- %s", bookName))
			newTitle = strings.ReplaceAll(newTitle, "· HonKit", fmt.Sprintf("- %s", bookName))
			newTitle = strings.ReplaceAll(newTitle, "&#xB7; GitBook", fmt.Sprintf("- %s", bookName))
			newTitle = strings.ReplaceAll(newTitle, "&#xB7; HonKit", fmt.Sprintf("- %s", bookName))

			// ==== remove the "A.2"-ish from the title
			titleParts = strings.Split(newTitle, " ")
			if strings.Contains(titleParts[0], ".") {
				titleParts = titleParts[1:]
			}
			newTitle = strings.Join(titleParts, " ")
		}
		htmlString = strings.ReplaceAll(htmlString, oldTitle, newTitle)

		// ==== adjust meta for SEO purpose
		metaReplacement := ""
		if isLandingPage {
			metaReplacement = `<meta content="Tutorial Gratis Belajar Dasar Pemrograman Golang Mulai Dari 0" name="description">`
		}
		htmlString = strings.ReplaceAll(htmlString, `<meta name="description" content="">`, metaReplacement)

		// ==== images' alt
		imagesAltToFind := ` alt="`
		imagesAltReplacement := imagesAltToFind + bookName + " - "
		htmlString = strings.ReplaceAll(htmlString, imagesAltToFind, imagesAltReplacement)

		// ==== disqus lazy load
		disqusJSBuf, err := os.ReadFile("./gitbook-plugin-disqus.js")
		if err != nil {
			fmt.Println("error", err)
		}
		err = os.WriteFile("./_book/gitbook/gitbook-plugin-disqus/plugin.js", disqusJSBuf, 0644)
		if err != nil {
			fmt.Println("error", err)
		}

		// ==== gitbook assets lazy load
		cssToLoad := []string{
			// "gitbook/style.css",
			"gitbook/gitbook-plugin-disqus/plugin.css",
			"gitbook/gitbook-plugin-highlight/website.css",
			"gitbook/gitbook-plugin-search/search.css",
			"gitbook/gitbook-plugin-fontsettings/website.css",
			// `/custom.css?v=` + getVersion() + `"`,
		}
		for _, cssFileNameToFind := range cssToLoad {
			cssFileNameReplacement := fmt.Sprintf(`%s" media="print" onload="this.media='all'`, cssFileNameToFind)
			htmlString = strings.ReplaceAll(htmlString, cssFileNameToFind, cssFileNameReplacement)
		}

		// ==== inject github stars button
		buttonToFind := `</body>`
		buttonReplacement := `<div style="position: fixed; top: 10px; right: 30px; padding: 10px; background-color: rgba(255, 255, 255, 0.7);">
			<a class="github-button" href="https://github.com/sponsors/novalagung" data-color-scheme="no-preference: light; light: light; dark: dark;" data-icon="octicon-heart" data-size="large" aria-label="Sponsor @novalagung on GitHub">Sponsor</a>&nbsp;
			<a class="github-button" href="https://github.com/novalagung/dasarpemrogramangolang" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star novalagung/dasarpemrogramangolang on GitHub">Star</a>&nbsp;
			<a class="github-button" href="https://github.com/novalagung" data-size="large" aria-label="Follow @novalagung on GitHub">Follow @novalagung</a>
			<script async defer src="https://buttons.github.io/buttons.js"></script>
		</div>` + buttonToFind
		htmlString = strings.ReplaceAll(htmlString, buttonToFind, buttonReplacement)

		// ==== inject adjustment css
		adjustmentCSSBuf, _ := os.ReadFile("./custom.css")
		os.WriteFile("./_book/gitbook/custom.css", adjustmentCSSBuf, 0644)
		adjustmentCSSToFind := `</head>`
		adjustmentCSSReplacement := `<link rel="stylesheet" href="gitbook/custom.css?v=` + getVersion() + `">` + adjustmentCSSToFind
		htmlString = strings.ReplaceAll(htmlString, adjustmentCSSToFind, adjustmentCSSReplacement)

		// ==== inject github stars js script
		buttonScriptToFind := `</head>`
		buttonScriptReplacement := `<script async defer src="https://buttons.github.io/buttons.js"></script>` + buttonScriptToFind
		htmlString = strings.ReplaceAll(htmlString, buttonScriptToFind, buttonScriptReplacement)

		// ==== inject ga4
		ga4propertyToFind := `</head>`
		ga4propertyReplacement := `<script async src="https://www.googletagmanager.com/gtag/js?id=` + ga4tagId + `"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());
			gtag('config', '` + ga4tagId + `');
		</script>` + ga4propertyToFind
		htmlString = strings.ReplaceAll(htmlString, ga4propertyToFind, ga4propertyReplacement)

		// ===== inject fb pixel
		// fbPixelToFind := `</head>`
		// fbPixelReplacement := `<script>!function(f,b,e,v,n,t,s){if(f.fbq)return;n=f.fbq=function(){n.callMethod?n.callMethod.apply(n,arguments):n.queue.push(arguments)};if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';n.queue=[];t=b.createElement(e);t.async=!0;t.src=v;s=b.getElementsByTagName(e)[0];s.parentNode.insertBefore(t,s)}(window,document,'script','https://connect.facebook.net/en_US/fbevents.js');fbq('init','1247398778924723');fbq('track','PageView');</script><noscript><imgheight="1"width="1"style="display:none"src="https://www.facebook.com/tr?id=1247398778924723&ev=PageView&noscript=1"/></noscript>` + fbPixelToFind
		// htmlString = strings.Replace(htmlString, fbPixelToFind, fbPixelReplacement)

		// ==== update file
		err = os.WriteFile(path, []byte(strings.TrimSpace(htmlString)), info.Mode())
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
	buf, err := os.ReadFile(siteMapPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	sitemapContent := string(buf)

	// ===== inject files into sitemap
	sitemapContent = strings.ReplaceAll(sitemapContent, `</urlset>`, strings.TrimSpace(`
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

	err = os.WriteFile(siteMapPath, []byte(sitemapContent), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("  ==>", siteMapPath)
}

func ebookPreAdjustment() {
	basePath, _ := os.Getwd()
	readmePath := filepath.Join(basePath, "README.md")

	buf, err := os.ReadFile(readmePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	mdString := string(buf)

	// ==== adjust version
	versionToFind := `((VERSION))`
	mdString = strings.ReplaceAll(mdString, versionToFind, getVersion())

	err = os.WriteFile(readmePath, []byte(mdString), 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ==== adjust content
	err = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".md" {
			return nil
		}

		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		htmlString := string(buf)

		// ==== remove substack embed
		substackEmbedToRemove := `<iframe src="partial/substack.html" width="100%" height="320px" frameborder="0" scrolling="no"></iframe>`
		htmlString = strings.ReplaceAll(htmlString, substackEmbedToRemove, "")

		// ==== remove ebooks embed
		ebooksEmbedToRemove := `<iframe src="partial/ebooks.html" width="100%" height="390px" frameborder="0" scrolling="no"></iframe>`
		htmlString = strings.ReplaceAll(htmlString, ebooksEmbedToRemove, "")

		// ==== update file
		err = os.WriteFile(path, []byte(strings.TrimSpace(htmlString)), info.Mode())
		if err != nil {
			return err
		}

		fmt.Println("  ==>", path)
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

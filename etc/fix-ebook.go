package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	flagAdjustment := flag.String("type", "", "adjustment type (pre/post)")
	flag.Parse()

	switch *flagAdjustment {
	case "pre":
		preAdjustment()
	default:
		break
	}
}

func preAdjustment() {
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
		substackEmbedToRemove := `<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>`
		htmlString = strings.ReplaceAll(htmlString, substackEmbedToRemove, "")

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

func getVersion() string {
	return fmt.Sprintf("%d.%s", baseVersion, now.Format("2006.01.02.150405"))
}

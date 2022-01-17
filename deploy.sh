#!/bin/sh

# cleanup
rm -rf book.json
git checkout README.md

# pull latest from upstream
git pull origin

# build website
cp book.json.template book.json
go run adjustment.go -type=pre
gitbook install
gitbook build
go run adjustment.go -type=post
cd _book
rm -rf LICENSE
rm -rf book*
rm -rf .git
rm -rf .gitignore
rm -rf .github
rm -rf *.md
rm -rf *.sh
rm -rf *.psd
rm -rf *.go
rm -rf *.css
rm -rf *.js

# build ebook files
cd ../
gitbook pdf ./ ./_book/dasarpemrogramangolang.pdf
gitbook epub ./ ./_book/dasarpemrogramangolang.epub
gitbook mobi ./ ./_book/dasarpemrogramangolang.mobi

echo "rebuild website complete ..."

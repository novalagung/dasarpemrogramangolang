mv book.json.temp book.json
mv _book .git _book_git
gitbook build
mv _book_git _book/.git
go run file-title-renamer.go -name "Dasar Pemrograman Golang"
cd _book
cp ./../dasarpemrogramangolang.pdf ./
cp ./../dasarpemrogramangolang.epub ./
cp ./../dasarpemrogramangolang.mobi ./
echo 'dasarpemrogramangolang.novalagung.com' > CNAME
rm -rf .git
rm -rf .gitignore
rm -rf *.psd
rm -rf *.sh
rm -rf *.go
rm -rf *.md
git init .
git add .
git commit -m "deploy"
git remote add github git@github.com:novalagung/dasarpemrogramangolang.git
git push -f github master:gh-pages
cd ..

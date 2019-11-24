mv book.json.temp book.json
mv _book/.git _book_git
gitbook build
mv _book_git _book/.git
rm -rf _book/_book_git
go run book-injector.go
cd _book
cp ./../dasarpemrogramangolang.pdf ./
cp ./../dasarpemrogramangolang.epub ./
cp ./../dasarpemrogramangolang.mobi ./
echo 'dasarpemrogramangolang.novalagung.com' > CNAME
echo 'google.com, pub-1417781814120840, DIRECT, f08c47fec0942fa0' > ads.txt
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

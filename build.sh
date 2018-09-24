gitbook build
cd _book
cp ./../pemrogramanwebgolang.pdf ./
cp ./../pemrogramanwebgolang.epub ./
cp ./../pemrogramanwebgolang.mobi ./
echo 'dasarpemrogramangolang.novalagung.com' > CNAME
rm -rf .git
rm -rf .gitignore
rm -rf *.psd
git init .
git add .
git commit -m "deploy"
git remote add github git@github.com:novalagung/dasarpemrogramangolang.git
git push -f github master:gh-pages
cd ..

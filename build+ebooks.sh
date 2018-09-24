gitbook build
gitbook pdf ./ ./dasarpemrogramangolang.pdf
gitbook epub ./ ./dasarpemrogramangolang.epub
gitbook mobi ./ ./dasarpemrogramangolang.mobi

cd _book
cp ./../dasarpemrogramangolang.pdf ./
cp ./../dasarpemrogramangolang.epub ./
cp ./../dasarpemrogramangolang.mobi ./
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

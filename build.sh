gitbook build
gitbook pdf ./ ./_book/dasarpemrogramangolang.pdf
gitbook epub ./ ./_book/dasarpemrogramangolang.epub
gitbook mobi ./ ./_book/dasarpemrogramangolang.mobi
cd _book
echo 'dasarpemrogramangolang.novalagung.com' > CNAME
rm -rf .git
rm -rf .gitignore
git init .
git add .
git commit -m "deploy"
git remote add github git@github.com:novalagung/dasarpemrogramangolang.git
git push -f github master:gh-pages


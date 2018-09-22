gitbook build
cd _book
echo 'dasarpemrogramangolang.novalagung.com' > CNAME
rm -rf .git
git init .
git add .
git commit -m "deploy"
git remote add github git@github.com:novalagung/dasarpemrogramangolang.git
git push -f github master:gh-pages


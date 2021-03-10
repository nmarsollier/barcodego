npm install apidoc
npm install apidoc-markdown2
rm -rf www
rm README-API.md
./node_modules/.bin/apidoc -i ./rest -o ./www
./node_modules/.bin/apidoc-markdown2 -p ./www -o README.md

echo "package cmd" > "cmd/version.go"
echo "var ZarubaVersion = \"$(git describe --tags --always)\"" >> "cmd/version.go"
git add . -A
git commit -m 'build'
go build -o zaruba
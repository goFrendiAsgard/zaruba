echo "package cmd" > "cmd/version.go"
echo "var ZarubaVersion = \"$(git describe --tags --always)\"" >> "cmd/version.go"

go build -o zaruba
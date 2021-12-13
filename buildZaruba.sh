# echo "package cmd" > "cmd/version.go"
# echo "var ZarubaVersion = \"$(git describe --tags --always)\"" >> "cmd/version.go"
go build -ldflags="-X 'github.com/state-alchemists/zaruba/cmd.ZarubaVersion=$(git describe --tags --always)'" -o zaruba
git add . -A
git commit -m 'build'
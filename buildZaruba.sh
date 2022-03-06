TAG=$(git describe $(git rev-list --tags --max-count=1))
BUILD=$(git rev-parse HEAD)
go build -ldflags="-X 'github.com/state-alchemists/zaruba/cmd.ZarubaVersion=${TAG}-${BUILD}'" -o zaruba
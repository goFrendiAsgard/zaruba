build:
	go build
	./build.sh
test:
	rm -Rf playground/projects/test-create
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
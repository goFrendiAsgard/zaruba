build:
	go build
test:
	mkdir -p ./coverage
	go test -v ./... --race -coverprofile=./coverage/profile.out -covermode=atomic
	go tool cover -html=./coverage/profile.out -o ./coverage/coverage.html
build:
	go build
test:
	mkdir -p ./coverage
	rm -f ./test_resource/alchemy/alembic.txt
	go test -v ./... --race -coverprofile=./coverage/profile.out -covermode=atomic
	go tool cover -html=./coverage/profile.out -o ./coverage/coverage.html
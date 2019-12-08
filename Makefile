build:
	go build
	./build.sh
test:
	set PATH=${PATH};$(pwd)
	mkdir -p "test-playground"
	# tear down
	# prepare
	# test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
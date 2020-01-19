export PATH := $(shell pwd):$(PATH)

build:
	./build.sh
install:
	./install.sh
test:
	./prepare-test.sh
	# perform test
	go test ./... 
test-verbose:
	./prepare-test.sh
	# perform test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
export PATH := $(shell pwd):$(PATH)

build:
	./build.sh
tear-down-test:
	./tear-down-test.sh
test:
	# build newest binary for current architecture
	go build
	# tear-down
	./tear-down-test.sh
	# prepare
	./prepare-test.sh
	# perform test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
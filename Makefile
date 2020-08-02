export PATH := $(shell pwd):$(PATH)

build:
	./build.sh
install:
	./install.sh --y
test:
	./prepare-test.sh
	# perform test
	go test -race ./... 
test-verbose:
	./prepare-test.sh
	# perform test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out -o coverage.html
test-integration:
	docker-compose -f docker-compose-testing.yml up --abort-on-container-exit
	docker-compose -f docker-compose-testing.yml down

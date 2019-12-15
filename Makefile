export PATH := $(shell pwd):$(PATH)

build:
	./build.sh
test:
	# build newest binary for current architecture
	go build
	# tear down
	rm -Rf ${ZARUBA_TEST_DIR}
	rm -Rf ${ZARUBA_TEMPLATE_DIR}
	# perform test
	go test ./...
test-verbose:
	# build newest binary for current architecture
	go build
	# tear down
	rm -Rf ${ZARUBA_TEST_DIR}
	rm -Rf ${ZARUBA_TEMPLATE_DIR}
	# perform test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
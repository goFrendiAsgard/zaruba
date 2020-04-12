export PATH := $(shell pwd):$(PATH)

build:
	./build.sh
install:
	./install.sh --y
test:
	./prepare-test.sh
	# perform test
	go test ./... 
test-verbose:
	./prepare-test.sh
	# perform test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out -o coverage.html
test-live:
	./prepare-test.sh
	./install.sh --y
	PATH="$(go env GOPATH)/bin:${PATH}" && \
	ZARUBA_TEMPLATE_DIR=${HOME}/.zaruba/template && \
	mkdir -p ${ZARUBA_TEST_DIR} && \
	cd ${ZARUBA_TEST_DIR} && \
	mkdir -p myproject && rm -Rf myproject && \
	zaruba create project myproject && \
	cd myproject && \
	zaruba create go-service . kodok && \
	zaruba create nodejs-service . siput && \
	zaruba run
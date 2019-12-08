export PATH := $(shell pwd):$(PATH)

build:
	go build
	./build.sh
test:
	go build

	# tear down
	rm -Rf templates/testInstallTemplate
	rm -Rf test-playground/testCreateComponent
	rm -Rf test-playground/testDo
	rm -Rf test-playground/testLink
	rm -Rf test-playground/testOrganize

	# prepare testLink
	mkdir -p test-playground/testLink

	# prepare testDo
	cp -r test-playground/testDo.template test-playground/testDo

	# prepare testOrganize
	cp -r test-playground/testOrganize.template test-playground/testOrganize

	# perform test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
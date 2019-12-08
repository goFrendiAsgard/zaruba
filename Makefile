build:
	go build
	./build.sh
test:
	# tear down
	rm -Rf templates/testInstallTemplate
	rm -Rf "test-playground"
	# prepare
	set PATH=${PATH};$(pwd)
	mkdir -p "test-playground"
	mkdir -p "test-playground/test-link"
	# test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
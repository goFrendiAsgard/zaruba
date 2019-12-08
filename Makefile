build:
	go build
	./build.sh
test:
	# tear down
	rm -Rf templates/project
	rm -Rf templates/templateForComponentTest
	rm -Rf "test-playground"
	# prepare
	set PATH=${PATH};$(pwd)
	mkdir -p "test-playground"
	# test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
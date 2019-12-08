build:
	go build
	./build.sh
test:
	# tear down
	rm -Rf templates/testInstallTemplate
	rm -Rf test-playground
	# prepare environment
	set PATH=${PATH};$(pwd)
	mkdir -p test-playground
	# prepare testLink
	mkdir -p test-playground/testLink
	# prepare testDo
	mkdir -p test-playground/testDo
	mkdir -p test-playground/testDo/subdir
	echo "#!/bin/sh\necho alpha > a.txt" > test-playground/testDo/doWrite
	echo "#!/bin/sh\necho beta > b.txt" > test-playground/testDo/subdir/doWrite
	# test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out
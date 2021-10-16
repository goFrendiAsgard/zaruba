build:
	go build
test:
	go build
	sudo rm -Rf playground
	./zaruba please clearLog
	./zaruba please test makeDocs
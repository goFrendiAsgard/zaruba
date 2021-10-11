build:
	go build
test:
	go build
	./zaruba please clearLog
	./zaruba please test makeDocs
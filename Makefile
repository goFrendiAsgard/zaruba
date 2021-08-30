build:
	go build
test:
	go build
	./zaruba please clearLog
	sudo rm -Rf playground
	./zaruba please test runPlaygroundDelay=10 makeTaskDocs
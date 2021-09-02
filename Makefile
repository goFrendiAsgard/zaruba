build:
	go build
test:
	go build
	./zaruba please clearLog
	./zaruba please test runPlaygroundDelay=10 makeTaskDocs
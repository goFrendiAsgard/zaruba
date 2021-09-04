build:
	go build
test:
	go build
	sudo rm -Rf playground
	mkdir -p playground
	./zaruba please clearLog
	./zaruba please test runPlaygroundDelay=10 makeTaskDocs
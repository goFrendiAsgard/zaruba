build:
	go build
test:
	go build
	sudo -E ./zaruba please setupUbuntu
	./zaruba please test
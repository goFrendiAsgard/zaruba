build:
	go build
test:
	go build
	sudo -E ./zaruba please setupUbuntu setup.injectBootstrap=false
	./zaruba please test setup.injectBootstrap=false runPlaygroundDelay=20
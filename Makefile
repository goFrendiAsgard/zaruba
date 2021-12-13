build:
	. ./buildZaruba.sh
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	./zaruba please clearLog
	./zaruba please test makeDocs
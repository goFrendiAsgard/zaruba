build:
	. ./buildZaruba.sh
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	./zaruba please clearLog
	./zaruba please test
adjustDocs:
	. ./buildZaruba.sh
	./zaruba please makeDocs
	cd docs && . ./.adjustDocs.sh	
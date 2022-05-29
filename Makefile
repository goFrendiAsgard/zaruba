build:
	. ./buildZaruba.sh
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	sudo rm -Rf docs/examples/playground
	./zaruba please test
adjustDocs:
	. ./buildZaruba.sh
	sudo rm -Rf docs/examples/playground
	./zaruba please makeDocs
	cd docs && . ./.adjustDocs.sh	
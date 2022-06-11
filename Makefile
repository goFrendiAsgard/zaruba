build:
	. ./buildZaruba.sh
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	sudo rm -Rf docs/examples/playground
	./zaruba please test
documentation:
	. ./buildZaruba.sh
	sudo rm -Rf docs/examples/playground
	./zaruba please makeDocs
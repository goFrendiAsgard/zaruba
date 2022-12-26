build:
	. ./buildZaruba.sh
test-toc:
	. ./buildZaruba.sh
	sudo rm -Rf test-resources/toc/playground
	cp -R test-resources/toc/template test-resources/toc/playground
	go test -v ./toc --race 
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	sudo rm -Rf docs/examples/playground
	./zaruba please test
documentation:
	. ./buildZaruba.sh
	sudo rm -Rf docs/examples/playground
	./zaruba please makeDocs
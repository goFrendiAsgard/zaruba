build:
	. ./buildZaruba.sh
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	./zaruba please clearLog
	./zaruba please test makeDocs
testFastApi:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	./zaruba please clearLog
	./zaruba please makeDocs testAddComponents
	cd playground && ../zaruba please removeProvoKafkaContainer removeProvoRabbitmqContainer
	cd playground && ../zaruba please startProvoKafka startProvoRabbitmq -t && ../zaruba please testProvoFastApi -e PROVO_FAST_API_TEST_INTEGRATION=1
	cd playground && ../zaruba please stopProvoKafkaContainer stopProvoRabbitmqContainer
adjustDocs:
	cd docs && . ./.adjustDocs.sh	
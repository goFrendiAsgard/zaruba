build:
	. ./buildZaruba.sh
test:
	. ./buildZaruba.sh
	sudo rm -Rf playground
	./zaruba please clearLog
	./zaruba please test makeDocs
adjustDocs:
	if [ -f "${HOME}/tocer/tocer.py" ]
	then
		cd docs
		echo "Run tocer"
		python "${HOME}/tocer/tocer.py" README.md 'alias zaruba="{{ .ZarubaBin }}"'
	fi
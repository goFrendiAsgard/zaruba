build:
	go build
test:
	mkdir -p ./coverage
	rm -f ./test_resource/alchemy/alembic.txt
	go test -v ./... --race -coverprofile=./coverage/profile.out -covermode=atomic
	go tool cover -html=./coverage/profile.out -o ./coverage/coverage.html
showOff:
	go build
	ZARUBA_HOME=$(realpath $(pwd))
	sudo rm -Rf playground
	mkdir -p playground
	./zaruba please clearLog
	echo "=== SORRY"
	./zaruba sorry
	echo "=== THANKS"
	./zaruba thanks
	echo "=== SETUP UBUNTU"
	cd playground && sudo -E ../zaruba please setupUbuntu
	echo "=== INIT PROJECT"
	cd playground && ../zaruba please initProject
	echo "=== ADD SUBREPOS"
	cd playground && ../zaruba please addSubrepo url="https://github.com/therealvasanth/fibonacci-clock" prefix="fibo"
	cd playground && ../zaruba please initSubrepos
	cd playground && ../zaruba please pullSubrepos
	echo "=== ADD FIBO SERVICE"
	cd playground && ../zaruba please addService location=fibo
	echo "=== ADD DOCKER SERVICE"
	cd playground && ../zaruba please addDocker image=rabbitmq
	echo "=== CREATE FASTAPI SERVICE"
	cd playground && ../zaruba please createService type=fastapi
	echo "=== RUN AND AUTOSTOP"
	cd playground && ../zaruba please run autostop
	echo "=== SHOW FASTAPI LOG"
	cd playground && ../zaruba please showLog task=fastapi
	echo "=== CLEAR LOG"
	cd playground && ../zaruba please clearLog
	echo "=== EXPLAIN START"
	cd playground && ../zaruba please explain start
	echo "=== DONE!!!"
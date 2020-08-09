export PATH := $(shell pwd):$(PATH)

build:
	./build.sh
install:
	./install.sh --y
test:
	docker-compose -f docker-compose-testing.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose-testing.yml down
publish-zaruba:
	docker build -f Dockerfile -t stalchmst/zaruba:latest .
	docker push stalchmst/zaruba:latest
publish-zaruba-env:
	# login as "stalchmst" first, docker -u <user> -p <password>
	docker build -f zaruba-env.Dockerfile -t stalchmst/zaruba-env:latest .
	docker push stalchmst/zaruba-env:latest


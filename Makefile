build:
	go build
test:
	mkdir -p ./coverage
	go test -v ./... --race -coverprofile=./coverage/profile.out -covermode=atomic
	go tool cover -html=./coverage/profile.out -o ./coverage/coverage.html
build-docker:
	docker build -t stalchmst/zaruba:latest .
run-docker:
	(docker stop zaruba || echo "container not running") && \
	(docker rm zaruba || echo "container not exist") && \
	docker run --name zaruba -p 2810:8080 -d stalchmst/zaruba:latest && \
	docker exec -it zaruba /bin/bash
publish-docker:
	docker push stalchmst/zaruba:latest

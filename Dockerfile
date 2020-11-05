FROM golang:1.14

ENV PATH="${PATH}:/.zaruba"
ENV DOCKER_HOST="tcp://host.docker.internal:2375"

RUN mkdir -p /.zaruba
COPY go.mod go.sum /.zaruba/

# build zaruba binary
RUN cd /.zaruba && go mod download
COPY . /.zaruba/
RUN cd /.zaruba && go build
RUN cd /.zaruba

RUN mkdir -p /project
WORKDIR /project

CMD "/bin/bash"

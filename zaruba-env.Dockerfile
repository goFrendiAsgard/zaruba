FROM ubuntu

ARG GO_DIR="/usr/local/go/bin"

ENV PATH="${GO_DIR}:${PATH}"
ENV DOCKER_HOST="tcp://host.docker.internal:2375"

RUN apt-get update && apt-get upgrade

RUN apt-get install make wget curl python3 nodejs npm golang docker.io -y
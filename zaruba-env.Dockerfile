FROM frolvlad/alpine-miniconda3

ARG GO_DIR="/usr/local/go/bin"

ENV PATH ${GO_DIR}:${PATH}
ENV DOCKER_HOST=tcp://host.docker.internal:2375

USER root

# install docker, git, nodejs, and npm
RUN apk update
RUN apk add --no-cache --allow-untrusted make gcc wget git docker musl-dev nodejs nodejs-npm go
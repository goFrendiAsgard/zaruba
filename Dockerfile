FROM golang:latest

ENV PATH ${PATH}:/app
ENV ZARUBA_SHELL /bin/bash
ENV ZARUBA_SHELL_ARG -c
ENV ZARUBA_TEMPLATE_DIR /app/templates

WORKDIR /app

# install docker and git
RUN apt-get update && apt-get install git docker.io -y

COPY . .

# build zaruba binary
RUN go mod download && go build -o zaruba .

CMD "./zaruba"

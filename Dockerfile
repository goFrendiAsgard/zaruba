FROM ubuntu:20.04

# preparing environments

ENV PATH="${PATH}:/.zaruba"
ENV DOCKER_HOST="tcp://host.docker.internal:2375"
ENV ZARUBA_HTTP_PORT=8080
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y \
    ncat make wget curl git \
    golang docker.io \
    python3 python3-pip python-is-python3 \
    nodejs npm 
RUN apt-get clean

RUN npm install -g typescript
RUN pip3 install pipenv

# building zaruba

RUN mkdir -p /.zaruba
COPY go.mod go.sum /.zaruba/

# build zaruba binary
RUN cd /.zaruba && go mod download
COPY . /.zaruba/
RUN cd /.zaruba && go build
RUN cd /.zaruba

RUN mkdir -p /project
WORKDIR /project

EXPOSE 8080

CMD ["sh", "-c", "zaruba please serveHttp port=${ZARUBA_HTTP_PORT}"]

FROM ubuntu:20.04

# preparing environments

ENV DOCKER_HOST="tcp://host.docker.internal:2375"
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
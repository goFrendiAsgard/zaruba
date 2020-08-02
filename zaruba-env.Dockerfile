FROM continuumio/anaconda3:latest

ENV PATH /usr/local/go/bin:${PATH}
ENV DOCKER_HOST=tcp://host.docker.internal:2375

# install docker, git, npm, and anaconda
RUN apt-get update --fix-missing && apt-get install -y make build-essential \
    git docker.io curl wget software-properties-common

# install latest node-js
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get install nodejs

# install typescript
RUN npm install -g typescript

# install latest golang
RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
RUN rm go1.13.linux-amd64.tar.gz


CMD "echo 'Welcome to zaruba env'"

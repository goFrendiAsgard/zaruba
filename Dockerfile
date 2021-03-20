FROM stalchmst/devbox:latest

# preparing environments

ENV PATH="${PATH}:/.zaruba"
ENV HOST_DOCKER_INTERNAL="host.docker.internal"
ENV DOCKER_HOST="tcp://${HOST_DOCKER_INTERNAL}:2375"

# building zaruba

RUN mkdir -p /.zaruba
COPY go.mod go.sum /.zaruba/

# build zaruba binary
RUN cd /.zaruba && go mod download
COPY . /.zaruba/
RUN cd /.zaruba && go build

# setup ubuntu
RUN zaruba please setupUbuntu setup.allowRoot=true
RUN apt-get clean

# setup pyenv, nvm, and kube
RUN zaruba please setupPyenv setupNvm setupKubeClient

RUN mkdir -p /project
WORKDIR /project

EXPOSE 2810

CMD ["sleep", "infinity"]
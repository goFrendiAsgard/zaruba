FROM stalchmst/devbox:latest

# preparing environments

ENV PATH="${PATH}:/.zaruba"
ENV ZARUBA_HOST_DOCKER_INTERNAL="host.docker.internal"
ENV DOCKER_HOST="tcp://${ZARUBA_HOST_DOCKER_INTERNAL}:2375"

# building zaruba

RUN mkdir -p /.zaruba
COPY go.mod go.sum /.zaruba/

# build zaruba binary
RUN cd /.zaruba && go mod download
COPY . /.zaruba/
RUN cd /.zaruba && go build

RUN mkdir -p /project
WORKDIR /project

CMD ["sleep", "infinity"]
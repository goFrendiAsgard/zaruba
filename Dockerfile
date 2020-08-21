FROM stalchmst/zaruba-env:latest

ENV PATH="/zaruba-src:${PATH}"
ENV DOCKER_HOST="tcp://host.docker.internal:2375"

ENV ZARUBA_SHELL="/bin/bash"
ENV ZARUBA_SHELL_ARG="-c"
ENV ZARUBA_DOCKER_HOST="host.docker.internal"
ENV ZARUBA_TEMPLATE_DIR="/zaruba-src/templates"
ENV ZARUBA_TEST_DIR="/tmp/zaruba-test"

# copy zaruba-src
RUN mkdir -p /zaruba-src
COPY . /zaruba-src

# build zaruba binary
RUN cd /zaruba-src && go mod download && go build -v -o zaruba .

RUN mkdir -p /project
WORKDIR /project

CMD "/bin/bash"

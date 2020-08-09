FROM stalchmst/zaruba-env:latest

ENV PATH /zaruba-src:${PATH}
ENV DOCKER_HOST=tcp://host.docker.internal:2375

ENV ZARUBA_SHELL /bin/sh
ENV ZARUBA_SHELL_ARG -c
ENV ZARUBA_DOCKER_HOST host.docker.internal
ENV ZARUBA_TEMPLATE_DIR /zaruba-src/templates
ENV ZARUBA_TEST_DIR=/tmp/zaruba-test

WORKDIR /zaruba-src

COPY . .

# install typescript
RUN npm install -g typescript

# build zaruba binary
RUN go mod download && go build -v -o zaruba .

CMD "mkdir -p /project && cd /project && (zaruba run || zaruba)"

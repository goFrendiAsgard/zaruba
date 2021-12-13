FROM golang:1.16 AS builder

RUN mkdir -p /zaruba
WORKDIR /zaruba
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
COPY ./.git ./.git
RUN chmod 755 ./buildZaruba.sh &&. ./buildZaruba.sh
RUN chmod 755 ./zaruba-tasks -R
RUN chmod 755 ./setup/*.sh

ENV ZARUBA_HOME="/zaruba"

FROM stalchmst/devbox:latest

ENV PATH="${PATH}:/.zaruba"
ENV DOCKER_HOST="tcp://host.docker.internal:2375"

RUN mkdir -p /.zaruba
COPY --from=builder /zaruba/zaruba /.zaruba/zaruba
COPY --from=builder /zaruba/core.zaruba.yaml /.zaruba/core.zaruba.yaml
COPY --from=builder /zaruba/advertisement.yaml /.zaruba/advertisement.yaml
COPY --from=builder /zaruba/zaruba-tasks /.zaruba/zaruba-tasks
COPY --from=builder /zaruba/setup /.zaruba/setup
COPY --from=builder /zaruba/setup/templates/bash/init.sh /.zaruba/init.sh
RUN chmod 755 /.zaruba/setup/*.sh && \
    chmod 755 /.zaruba/init.sh && \
    mkdir -p /project

WORKDIR /project

CMD ["sleep", "infinity"]
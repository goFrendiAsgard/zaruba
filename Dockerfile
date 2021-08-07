FROM golang:1.16 AS builder

RUN mkdir -p /zaruba
WORKDIR /zaruba
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
COPY ./.git ./.git
RUN go build -o zaruba

ENV ZARUBA_HOME /zaruba
RUN chmod 755 ./bash/*.sh
RUN chmod 755 ./write_version.sh
RUN ./write_version.sh


FROM stalchmst/devbox:latest

ENV PATH="${PATH}:/.zaruba"
ENV ZARUBA_HOST_DOCKER_INTERNAL="host.docker.internal"
ENV DOCKER_HOST="tcp://${ZARUBA_HOST_DOCKER_INTERNAL}:2375"

RUN mkdir -p /.zaruba
COPY --from=builder /zaruba/zaruba ./.zaruba/zaruba
COPY --from=builder /zaruba/.version ./.zaruba/.version
COPY --from=builder /zaruba/advertisement.yaml ./.zaruba/advertisement.yaml
COPY --from=builder /zaruba/bash ./.zaruba/bash
COPY --from=builder /zaruba/scripts ./.zaruba/scripts
COPY --from=builder /zaruba/setup ./.zaruba/setup
COPY --from=builder /zaruba/templates ./.zaruba/templates

RUN mkdir -p /project
WORKDIR /project

CMD ["sleep", "infinity"]
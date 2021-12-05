FROM golang:1.16 AS builder

RUN mkdir -p /zaruba
WORKDIR /zaruba
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
COPY ./.git ./.git
RUN go build -o zaruba && \
    chmod 755 ./zaruba-tasks -R \
    chmod 755 ./setup/*.sh

ENV ZARUBA_HOME="/zaruba"
RUN . ./bash/util.sh && getVersion > /zaruba/.version

FROM stalchmst/devbox:latest

ENV PATH="${PATH}:/.zaruba"
ENV DOCKER_HOST="tcp://host.docker.internal:2375"

RUN mkdir -p /.zaruba
COPY --from=builder /zaruba/zaruba /.zaruba/zaruba
COPY --from=builder /zaruba/.version /.zaruba/.version
COPY --from=builder /zaruba/advertisement.yaml /.zaruba/advertisement.yaml
COPY --from=builder /zaruba/bash /.zaruba/bash
COPY --from=builder /zaruba/scripts /.zaruba/scripts
COPY --from=builder /zaruba/setup /.zaruba/setup
COPY --from=builder /zaruba/templates/bash/init.sh /.zaruba/init.sh
COPY --from=builder /zaruba/templates /.zaruba/templates
RUN chmod 755 /.zaruba/setup/*.sh && \
    chmod 755 /.zaruba/init.sh && \
    mkdir -p /project

WORKDIR /project

CMD ["sleep", "infinity"]
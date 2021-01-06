#!/bin/sh

# USAGE
# sh inspect_docker.sh <object> <format> <container-name>
#
# EXAMPLE
# sh inspect_docker.sh container '.State.Running' 'rmq'

_OBJECT="${1}"
_FORMAT="${2}"
_CONTAINER_NAME="${3}"

set +e
docker ${_OBJECT} inspect -f "{{ ${_FORMAT} }}" "${_CONTAINER_NAME}"

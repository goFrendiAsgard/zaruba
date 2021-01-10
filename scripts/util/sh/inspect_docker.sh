# USAGE
# sh inspect_docker.sh <object> <format> <container-name>
#
# EXAMPLE
# sh inspect_docker.sh container '.State.Running' 'rmq'

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

_OBJECT="${1}"
_FORMAT="${2}"
_CONTAINER_NAME="${3}"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e
docker ${_OBJECT} inspect -f "{{ ${_FORMAT} }}" "${_CONTAINER_NAME}"

set "${_OLD_STATE}"
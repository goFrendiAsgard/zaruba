# USAGE
# sh stop_container.sh <container-name>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

docker image inspect "${1}" > /dev/null || docker  pull "${1}"

set "${_OLD_STATE}"
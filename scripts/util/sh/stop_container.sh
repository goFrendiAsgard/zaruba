# USAGE
# sh stop_container.sh <container-name>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

if [ "$(docker container inspect -f "{{ .State.Running }}" "${1}")" = true ]
then
    docker stop "${1}"
fi

set "${_OLD_STATE}"
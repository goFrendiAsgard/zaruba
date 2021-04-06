# USAGE
# sh remove_container.sh <container-name>

. "${ZARUBA_HOME}/scripts/util/sh/_include.sh"

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

if [ ! -z $(docker container inspect -f "{{ .Name }}" "${1}") ]
then
    docker rm "${1}"
fi

set "${_OLD_STATE}"
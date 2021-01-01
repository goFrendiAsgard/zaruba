#!/bin/sh

# USAGE
# sh stop_and_remove_container.sh <container-name>
#

(echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
set +e

if [ "$(docker container inspect -f "{{ .State.Running }}" "${1}")" = true ]
then
    docker stop "${1}"
fi
if [ ! -z $(docker container inspect -f "{{ .Name }}" "${1}") ]
then
    docker rm "${1}"
fi

set "${_OLD_STATE}"

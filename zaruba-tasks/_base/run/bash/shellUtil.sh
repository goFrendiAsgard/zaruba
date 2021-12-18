if [ "${ZARUBA_INIT_SCRIPT_LOADED}" != "1" ] 
then
    if [ -f "${ZARUBA_HOME}/init.sh" ]
    then
        . "${ZARUBA_HOME}/init.sh"
    elif [ -f "${ZARUBA_HOME}/setup/templates/bash/init.sh" ]
    then
        . "${ZARUBA_HOME}/setup/templates/bash/init.sh"
    fi
fi

# USAGE getValueOrDefault <value> <default>
getValueOrDefault() {
    if [ -z "${1}" ]
    then
        echo "${2}"
    else
        echo "${1}"
    fi
}

getLatestGitCommit() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git rev-parse --verify HEAD
    set "${_OLD_STATE}"
}


getLatestGitTagCommit() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git rev-parse --verify "$(git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags)"^{commit}
    set "${_OLD_STATE}"
}


getLatestGitTag() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags
    set "${_OLD_STATE}"
}

# USAGE: gitInitSubrepo <name> <prefix> <url> <branch>
gitInitSubrepo() {
    _NAME="${1}"
    _PREFIX="${2}"
    _URL="${3}"
    _BRANCH="${4}"
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git remote add "${_NAME}" "${_URL}"
    git subtree add --prefix="${_PREFIX}" "${_NAME}" "${_BRANCH}"
    git fetch "${_NAME}" "${_BRANCH}"
    git pull "${_NAME}" "${_BRANCH}"
    set "${_OLD_STATE}"
}


# USAGE: gitSave <message>
gitSave() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git add . -A
    git commit -m "💀 ${1}"
    set "${_OLD_STATE}"
}


# USAGE: inspectDocker <object> <format> <container-name>
inspectDocker() {
    _OBJECT_TYPE="${1}"
    _FORMAT="${2}"
    _OBJECT_NAME="${3}"
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    docker ${_OBJECT_TYPE} inspect -f "{{ ${_FORMAT} }}" "${_OBJECT_NAME}"
    set "${_OLD_STATE}"
}

# USAGE: isCommandError <command>
isCommandError() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    $@ >> /dev/null
    _STATUS=$?
    if [ "$_STATUS" = 0 ]
    then
        echo 0
    else
        echo 1
    fi 
    set "${_OLD_STATE}"
}


# USAGE: isCommandExist <command>
isCommandExist() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    $@ >> /dev/null
    _STATUS=$?
    if [ "$_STATUS" = 127 ]
    then
        echo 0
    else
        echo 1
    fi 
    set "${_OLD_STATE}"
}


# USAGE: linkResource <src> <dst>
linkResource() {
    _SRC="${1}"
    _DST="${2}"
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set -e
    if [ -e "${_DST}" ]
    then
        chmod 777 -R "${_DST}" && rm -Rf "${_DST}" && cp -rnT "${_SRC}" "${_DST}" && chmod 555 -R "${_DST}"
    fi
    cp -rnT "${_SRC}" "${_DST}" && chmod 555 -R "${_DST}"
    set "${_OLD_STATE}"
}


# USAGE: pullImage <image>
pullImage() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    docker image inspect "${1}" > /dev/null || docker  pull "${1}"
    set "${_OLD_STATE}"
} 


# USAGE: isContainerExist <container>
isContainerExist() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    __CONTAINER_NAME="$(docker container inspect -f "{{ .Name }}" "${1}")"
    if [ -z "${__CONTAINER_NAME}" ]
    then
        echo 0
    else
        echo 1
    fi
    set "${_OLD_STATE}"
} 


# USAGE: getContainerStatus <container>
getContainerStatus() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    if [ "$(isContainerExist "${1}")" = 1 ]
    then
        docker container inspect -f "{{ .State.Status }}" "${1}"
    fi
    set "${_OLD_STATE}"
}


# USAGE: removeContainer <container>
removeContainer() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    if [ "$(isContainerExist "${1}")" = 1 ]
    then
        docker rm "${1}"
    fi
    set "${_OLD_STATE}"
} 


# USAGE: stopContainer <container>
stopContainer() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    if [ "$(isContainerExist "${1}")" = 1 ] && [ "$(getContainerStatus "${1}" )" != "exited" ]
    then
        docker stop "${1}"
    fi
    set "${_OLD_STATE}"
} 


# USAGE: waitPort <host> <port>
waitPort() {
    until nc -z ${1} ${2}
    do
        sleep 1
    done
} 
get_latest_git_commit() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git rev-parse --verify HEAD
    set "${_OLD_STATE}"    
}


get_latest_git_tag_commit() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git rev-parse --verify "$(git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags)"^{commit}
    set "${_OLD_STATE}"
}


get_latest_git_tag() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git for-each-ref --sort=-taggerdate --count=1 --format '%(tag)' refs/tags
    set "${_OLD_STATE}"
}


show_version() {
    if [ -z "$(get_latest_git_tag)" ]
    then
        echo "Dev - $(get_latest_git_commit)"
    elif [ "$(get_latest_git_tag_commit)" = "$(get_latest_git_commit)" ]
    then
        echo "$(get_latest_git_tag) - $(get_latest_git_commit)"
    else
        echo "Dev - $(get_latest_git_tag) - $(get_latest_git_commit)"
    fi
}


# USAGE: git_init_subrepo <name> <prefix> <url> <branch>
git_init_subrepo() {
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


# USAGE: git_save <message>
git_save() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    git add . -A
    git commit -m "💀 ${1}"
    set "${_OLD_STATE}"
}


# USAGE: inspect_docker <object> <format> <container-name>
inspect_docker() {
    _OBJECT="${1}"
    _FORMAT="${2}"
    _CONTAINER_NAME="${3}"
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    docker ${_OBJECT} inspect -f "{{ ${_FORMAT} }}" "${_CONTAINER_NAME}"
    set "${_OLD_STATE}"
}


# USAGE: is_command_error <command>
is_command_error() {
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


# USAGE: is_command_exist <command>
is_command_exist() {
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


# USAGE: link_resource <src> <dst>
link_resource() {
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


# USAGE: pull_image <image>
pull_image() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    docker image inspect "${1}" > /dev/null || docker  pull "${1}"
    set "${_OLD_STATE}"
} 


# USAGE: remove_container <container>
remove_container() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    if [ ! -z $(docker container inspect -f "{{ .Name }}" "${1}") ]
    then
        docker rm "${1}"
    fi
    set "${_OLD_STATE}"
} 


# USAGE: stop_container <container>
stop_container() {
    (echo $- | grep -Eq ^.*e.*$) && _OLD_STATE=-e || _OLD_STATE=+e
    set +e
    if [ ! -z $(docker container inspect -f "{{ .Name }}" "${1}") ]
    then
        docker stop "${1}"
    fi
    set "${_OLD_STATE}"
} 


# USAGE: wait_port <host> <port>
wait_port() {
    until nc -z ${1} ${2}
    do
        sleep 1
    done
} 
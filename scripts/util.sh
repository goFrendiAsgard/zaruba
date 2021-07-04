Normal=$'\033[0m'
Bold=$'\033[1m'
Faint=$'\033[2m'
Italic=$'\033[3m'
Underline=$'\033[4m'
BlinkSlow=$'\033[5m'
BlinkRapid=$'\033[6m'
Inverse=$'\033[7m'
Conceal=$'\033[8m'
CrossedOut=$'\033[9m'
Black=$'\033[30m'
Red=$'\033[31m'
Green=$'\033[32m'
Yellow=$'\033[33m'
Blue=$'\033[34m'
Magenta=$'\033[35m'
Cyan=$'\033[36m'
White=$'\033[37m'
BgBlack=$'\033[40m'
BgRed=$'\033[41m'
BgGreen=$'\033[42m'
BgYellow=$'\033[43m'
BgBlue=$'\033[44m'
BgMagenta=$'\033[45m'
BgCyan=$'\033[46m'
BgWhite=$'\033[47m'
NoStyle=$'\033[0m'
NoUnderline=$'\033[24m'
NoInverse=$'\033[27m'
NoColor=$'\033[39m'

# USAGE create_docker_task <template_location> <image_name> <container_name> <service_name>
create_docker_task() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_docker_task.py" "${1}" "${2}" "${3}" "${4}" "${5}" "${6}"
}


# USAGE create_fast_crud <template_location> <service_name> <module_name> <entity_name> <field_names>
create_fast_crud() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_fast_crud.py" "${1}" "${2}" "${3}" "${4}" "${5}"
}


# USAGE create_fast_event_handler <template_location> <service_name> <module_name> <event_name>
create_fast_event_handler() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_fast_event_handler.py" "${1}" "${2}" "${3}" "${4}"
}


# USAGE create_fast_module <template_location> <service_name> <module_name>
create_fast_module() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_fast_module.py" "${1}" "${2}" "${3}"
}


# USAGE create_fast_route <template_location> <service_name> <module_name> <http_method> <url>
create_fast_route() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_fast_route.py" "${1}" "${2}" "${3}" "${4}" "${5}"
}


# USAGE create_fast_rpc_handler <template_location> <service_name> <module_name> <event_name>
create_fast_rpc_handler() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_fast_rpc_handler.py" "${1}" "${2}" "${3}" "${4}"
}


# USAGE create_fast_service <template_location> <service_name>
create_fast_service() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_fast_service.py" "${1}" "${2}"
}


# USAGE create_helm_deployment <service_name>
create_helm_deployment() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_helm_deployment.py" "${1}"
}


# USAGE create_helm_task
create_helm_task() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_helm_task.py" "${1}"
}


# USAGE create_service_task <template_location> <service_name> <image_name> <container_name> <location> <start_command> <ports> <runner_version>
create_service_task() {
    PIPENV_IGNORE_VIRTUAL_ENVS=1 PIPENV_DONT_LOAD_ENV=1 PIPENV_PIPFILE="${ZARUBA_HOME}/scripts/python/Pipfile" pipenv run python "${ZARUBA_HOME}/scripts/python/create_service_task.py" "${1}" "${2}" "${3}" "${4}" "${5}" "${6}" "${7}" "${8}" "${9}" "${10}"
}

# USAGE update_env <projectLocation>
update_env() {
    "${ZARUBA_HOME}/zaruba" updateProjectTaskEnv "${1}/main.zaruba.yaml"
    "${ZARUBA_HOME}/zaruba" updateProjectEnvFiles "${1}/main.zaruba.yaml"
}


# USAGE: append_if_exist <str> <file>
append_if_exist() {
    if [ -f "${2}" ]
    then
        echo "" >> "${2}"
        echo "${1}" >> "${2}"
    fi
}


get_current_user() {
    if [ ! -z "$SUDO_USER" ]
    then
        echo "$SUDO_USER"
    elif [ ! -z "$USER" ]
    then
        echo "$USER"
    else
        id -u -n
    fi
}


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


# USAGE: inject_bootstrap <bashrc-path>
inject_bootstrap() {
    if [ -f "${1}" ]
    then
        if cat "${1}" | grep -Fqe "${ZARUBA_HOME}/scripts/bootstrap.sh"
        then
            echo -e "${Faint}Bootstrap script ${ZARUBA_HOME}/scripts/bootstrap.sh is already loaded in ${1}${Normal}"
        else
            echo "" >> "${1}"
            echo "# Load zaruba's bootstrap" >> "${1}"
            echo "if [ -x "${ZARUBA_HOME}/scripts/bootstrap.sh" ]" >> "${1}"
            echo 'then' >> "${1}"
            echo "    . ${ZARUBA_HOME}/scripts/bootstrap.sh" >> "${1}"
            echo 'fi' >> "${1}"
            echo "" >> "${1}"
        fi
    fi 
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
    echo "${Yellow}Link \"${_SRC}\" into \"${_DST}\"${Normal}"
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


# USAGE: should_be_dir <value> <error-message>
should_be_dir() {
    if [ ! -d "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_be_file <value> <error-message>
should_be_file() {
    if [ ! -f "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_be_exist <value> <error-message>
should_be_exist() {
    if [ ! -e "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_be_empty <value> <error-message>
should_be_empty() {
    if [ ! -z "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_not_be_dir <value> <error-message>
should_not_be_dir() {
    if [ -d "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_not_be_file <value> <error-message>
should_not_be_file() {
    if [ -f "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_not_be_exist <value> <error-message>
should_not_be_exist() {
    if [ -e "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


# USAGE: should_not_be_empty <value> <error-message>
should_not_be_empty() {
    if [ -z "${1}" ]
    then
        echo "${2}" 1>&2
        exit 1
    fi
}


show_version() {
    cd ${ZARUBA_HOME}
    if [ -z "$(get_latest_git_tag)" ]
    then
        echo "Current version : Dev - $(get_latest_git_commit)"
    elif [ "$(get_latest_git_tag_commit)" = "$(get_latest_git_commit)" ]
    then
        echo "${Bold}${Yellow}Current version : $(get_latest_git_tag) - $(get_latest_git_commit)${Normal}"
    else
        echo "${Bold}${Yellow}Current version : Dev - $(get_latest_git_commit)${Normal}"
        echo "${Bold}${Yellow}Latest tag      : $(get_latest_git_tag)${Normal}"
    fi
}


get_current_user() {
    if [ ! -z "$SUDO_USER" ]
    then
        echo "$SUDO_USER"
    elif [ ! -z "$USER" ]
    then
        echo "$USER"
    else
        id -u -n
    fi
}


init_bootstrap() {
    _CURRENT_USER=$(get_current_user)
    _BOOTSTRAP_SCRIPT="${ZARUBA_HOME}/scripts/bootstrap.sh"
    if [ ! -f "${_BOOTSTRAP_SCRIPT}" ]
    then
        touch "${_BOOTSTRAP_SCRIPT}"
        chmod 755 "${_BOOTSTRAP_SCRIPT}"
        chown "${_CURRENT_USER}" "${_BOOTSTRAP_SCRIPT}"
    fi
    . "${_BOOTSTRAP_SCRIPT}"
    # also include .local/bin
    if echo "$PATH" | grep -Fqe "${HOME}/.local/bin"
    then
        echo "${Faint}${HOME}/.local/bin is already in the PATH${Normal}"
    else
        TEMPLATE_CONTENT="$(cat "${ZARUBA_HOME}/scripts/templates/shell/include_local_bin.sh")"
        append_if_exist "${TEMPLATE_CONTENT}" "${_BOOTSTRAP_SCRIPT}"
        . "${_BOOTSTRAP_SCRIPT}"
    fi    
}
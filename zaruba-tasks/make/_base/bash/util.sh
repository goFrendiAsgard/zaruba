_setReplacementMap() {
    __ZRB_KEY="${1}"
    __ZRB_VAL="${2}"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "${1}" "${2}")"
}

_readText() {
    __ZRB_FILE_NAME="${1}"
    "${ZARUBA_BIN}" file read "${__ZRB_FILE_NAME}"
}

_readLines() {
    __ZRB_FILE_NAME="${1}"
    "${ZARUBA_BIN}" lines read "${__ZRB_FILE_NAME}"
}

_indent() {
    __ZRB_LINE="${1}"
    __ZRB_INDENTATION="${2}"
    "${ZARUBA_BIN}" str fullIndent "${__ZRB_LINE}" "${__ZRB_INDENTATION}"
}

_getIndentation() {
    __ZRB_LINE="${1}"
    "${ZARUBA_BIN}" str getIndentation "${__ZRB_LINE}"
}

_getLineFromFile() {
    __ZRB_FILE_NAME="${1}"
    __ZRB_INDEX="${2}"
    "${ZARUBA_BIN}" file getLine "${__ZRB_FILE_NAME}" "${__ZRB_INDEX}"
}

_getSubmatchFromFile() {
    __ZRB_FILE_NAME="${1}"
    __ZRB_PATTERN="${2}"
    __ZRB_INDEX="${3}"
    if [ "$("${ZARUBA_BIN}" list validate "${__ZRB_PATTERN}")" = 0 ]
    then
        __ZRB_PATTERN="$("${ZARUBA_BIN}" list append "[]" "${__ZRB_PATTERN}")"
    fi
    if [ -z "${__ZRB_INDEX}" ]
    then
        __ZRB_INDEX="-1"
    fi
    "${ZARUBA_BIN}" file submatch "${__ZRB_FILE_NAME}" "${__ZRB_PATTERN}" --index="${__ZRB_INDEX}"
}

_getLineIndexFromFile() {
    __ZRB_FILE_NAME="${1}"
    __ZRB_PATTERN="${2}"
    if [ "$("${ZARUBA_BIN}" list validate "${__ZRB_PATTERN}")" = 0 ]
    then
        __ZRB_PATTERN="$("${ZARUBA_BIN}" list append "[]" "${__ZRB_PATTERN}")"
    fi
    "${ZARUBA_BIN}" file getLineIndex "${__ZRB_FILE_NAME}" "${__ZRB_PATTERN}"
}

_getPartialContent() {
    __ZRB_TEMPLATE_FILE_NAME="${1}"
    __ZRB_TEMPLATE="$(_readText "${__ZRB_TEMPLATE_FILE_NAME}")"
    __ZRB_NEW_CONTENT="$("${ZARUBA_BIN}" str replace "${__ZRB_TEMPLATE}" "${_ZRB_REPLACEMENT_MAP}" )"
    echo "${__ZRB_NEW_CONTENT}"
}

_insertPartialBefore() {
    __ZRB_FILE_NAME="${1}"
    __ZRB_CONTENT="${2}"
    __ZRB_INDEX="${3}"
    if [ -z "${__ZRB_INDEX}" ]
    then
        __ZRB_INDEX="0"
    fi
    "${ZARUBA_BIN}" file insertBefore "${__ZRB_FILE_NAME}" "${__ZRB_CONTENT}" --index="${__ZRB_INDEX}"
}

_insertPartialAfter() {
    __ZRB_FILE_NAME="${1}"
    __ZRB_CONTENT="${2}"
    __ZRB_INDEX="${3}"
    if [ -z "${__ZRB_INDEX}" ]
    then
        __ZRB_INDEX="-1"
    fi
    "${ZARUBA_BIN}" file insertAfter "${__ZRB_FILE_NAME}" "${__ZRB_CONTENT}" --index="${__ZRB_INDEX}"
}

_replacePartialAt() {
    __ZRB_FILE_NAME="${1}"
    __ZRB_CONTENT="${2}"
    __ZRB_INDEX="${3}"
    if [ -z "${__ZRB_INDEX}" ]
    then
        __ZRB_INDEX="0"
    fi
    "${ZARUBA_BIN}" file replace "${__ZRB_FILE_NAME}" "${__ZRB_CONTENT}" --index="${__ZRB_INDEX}"
}

_addConfigToReplacementMap() {
    # add config with prefix: 'ztplCfg'
    __ZRB_CONFIG_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map transformKey "${_ZRB_CONFIG_MAP}" -t pascal -p ztplCfg)"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map merge "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_CONFIG_REPLACEMENT_MAP}")"
}

_addEnvToReplacementMap() {
    # add env with prefix: 'ZTPL_ENV_'
    __ZRB_ENV_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map transformKey "${_ZRB_ENV_MAP}" -p ZTPL_ENV_)"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map merge "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_ENV_REPLACEMENT_MAP}")"
}

_addTaskDependency() {
    __ZRB_TASK_NAME="${1}"
    __ZRB_DEPENDENCY_TASK_NAME="${2}"
    __ZRB_CREATE_TASK="${3}"
    __ZRB_PROJECT_FILE_NAME="${4}"
    # check dependency task existance
    echo "Checking ${__ZRB_DEPENDENCY_TASK_NAME}"
    if [ "$("${ZARUBA_BIN}" task isExist "${__ZRB_DEPENDENCY_TASK_NAME}" "${__ZRB_PROJECT_FILE_NAME}")" = 1 ]
    then
        # check dependee task existance
        echo "Checking ${__ZRB_TASK_NAME}"
        if [ "${__ZRB_CREATE_TASK}" = 1 ]
        then
            "${ZARUBA_BIN}" project addTask "${__ZRB_TASK_NAME}" "${__ZRB_PROJECT_FILE_NAME}"
        elif [ "$("${ZARUBA_BIN}" task isExist "${__ZRB_TASK_NAME}" "${__ZRB_PROJECT_FILE_NAME}")" = 0 ]
        then
            echo "Task ${__ZRB_TASK_NAME} doesn't exist"
            return
        fi
        # link dependency task to task
        echo "Adding ${__ZRB_DEPENDENCY_TASK_NAME} as dependency of ${__ZRB_TASK_NAME}"
        "${ZARUBA_BIN}" task addDependencies "${__ZRB_TASK_NAME}" "[\"${__ZRB_DEPENDENCY_TASK_NAME}\"]" "${__ZRB_PROJECT_FILE_NAME}" 
    fi
}

_generate() {
    __ZRB_TEMPLATE_LOCATIONS="${1}"
    __ZRB_REPLACEMENT_MAP="${2}"
    for __ZRB_TEMPLATE_INDEX in $("${ZARUBA_BIN}" list rangeIndex "${__ZRB_TEMPLATE_LOCATIONS}")
    do 
        __ZRB_TEMPLATE_LOCATION="$("${ZARUBA_BIN}" list get "${__ZRB_TEMPLATE_LOCATIONS}" "${__ZRB_TEMPLATE_INDEX}")"
        ${ZARUBA_HOME}/zaruba generate "${__ZRB_TEMPLATE_LOCATION}" "." "${__ZRB_REPLACEMENT_MAP}"
    done
}

_getYamlEnvs() {
    __ZRB_ENVS="${1}"
    __ZRB_ENV_PREFIX="${2}"
    __ZRB_MAP_ENVS='{}'
    __OLD_IFS="${IFS}"
    IFS=$'\n'
    for __ZRB_KEY in $("${ZARUBA_BIN}" map rangeKey "${__ZRB_ENVS}")
    do
        __ZRB_FROM="${__ZRB_ENV_PREFIX}_${__ZRB_KEY}"
        __ZRB_DEFAULT="$("${ZARUBA_BIN}" map get "${__ZRB_ENVS}" "${__ZRB_KEY}")"
        __ZRB_DEFAULT="$("${ZARUBA_BIN}" str doubleQuote "${__ZRB_DEFAULT}")"
        __ZRB_SINGLE_MAP_ENV='{}'
        __ZRB_SINGLE_MAP_ENV="$("${ZARUBA_BIN}" map set "${__ZRB_SINGLE_MAP_ENV}" "from" "${__ZRB_FROM}")"
        __ZRB_SINGLE_MAP_ENV="$("${ZARUBA_BIN}" map set "${__ZRB_SINGLE_MAP_ENV}" "default" "${__ZRB_DEFAULT}")"
        __ZRB_MAP_ENVS="$("${ZARUBA_BIN}" map set "${__ZRB_MAP_ENVS}" "${__ZRB_KEY}" "${__ZRB_SINGLE_MAP_ENV}")"
    done
    IFS="${__OLD_IFS}"
    __ZRB_YAML_ENVS="$("${ZARUBA_BIN}" yaml print "${__ZRB_MAP_ENVS}")"
    echo "${__ZRB_YAML_ENVS}"
}

_skipIfPathExist() {
    __ZRB_PATH="${1}"
    if [ ! -z "${__ZRB_PATH}" ] && [ -e "${__ZRB_PATH}" ]
    then
        echo "${_YELLOW}[SKIP] ${__ZRB_PATH} already exist.${_NORMAL}"
        exit 0
    fi
}
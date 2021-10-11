_setReplacementMap() {
    __ZRB_KEY="${1}"
    __ZRB_VAL="${2}"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_REPLACEMENT_MAP}" "${__ZRB_KEY}" "${__ZRB_VAL}")"
}

_registerTask() {
    __ZRB_PROJECT_FILE_NAME="${1}"
    __ZRB_MAIN_TASK_NAME="${2}"
    __ZRB_MODULE_TASK_NAME="${3}"
    if [ "$("${ZARUBA_HOME}/zaruba" task isExist "${__ZRB_PROJECT_FILE_NAME}" "${__ZRB_MODULE_TASK_NAME}")" = 1 ]
    then
        if [ "$("${ZARUBA_HOME}/zaruba" task isExist "${__ZRB_PROJECT_FILE_NAME}" "${__ZRB_MAIN_TASK_NAME}")" = 0 ]
        then
            "${ZARUBA_HOME}/zaruba" project addTask "${__ZRB_PROJECT_FILE_NAME}" "${__ZRB_MAIN_TASK_NAME}"
        fi
        "${ZARUBA_HOME}/zaruba" task addDependency "${__ZRB_PROJECT_FILE_NAME}" "${__ZRB_MAIN_TASK_NAME}" "[\"${__ZRB_MODULE_TASK_NAME}\"]"
    fi
}

_generate() {
    __ZRB_TEMPLATE_LOCATIONS="${1}"
    __ZRB_REPLACEMENT_MAP="${2}"
    for __ZRB_TEMPLATE_INDEX in $("${ZARUBA_HOME}/zaruba" list rangeIndex "${__ZRB_TEMPLATE_LOCATIONS}")
    do 
        __ZRB_TEMPLATE_LOCATION="$("${ZARUBA_HOME}/zaruba" list get "${__ZRB_TEMPLATE_LOCATIONS}" "${__ZRB_TEMPLATE_INDEX}")"
        ${ZARUBA_HOME}/zaruba generate "${__ZRB_TEMPLATE_LOCATION}" "." "${__ZRB_REPLACEMENT_MAP}"
    done
}

_getYamlEnvs() {
    __ZRB_ENVS="${1}"
    __ZRB_ENV_PREFIX="${2}"
    __ZRB_MAP_ENVS='{}'
    for __ZRB_KEY in $("${ZARUBA_HOME}/zaruba" map rangeKey "${__ZRB_ENVS}")
    do
        __ZRB_FROM="${__ZRB_ENV_PREFIX}_${__ZRB_KEY}"
        __ZRB_DEFAULT="$("${ZARUBA_HOME}/zaruba" map get "${__ZRB_ENVS}" "${__ZRB_KEY}")"
        __ZRB_SINGLE_MAP_ENV='{}'
        __ZRB_SINGLE_MAP_ENV="$("${ZARUBA_HOME}/zaruba" map set "${__ZRB_SINGLE_MAP_ENV}" "from" "${__ZRB_FROM}")"
        __ZRB_SINGLE_MAP_ENV="$("${ZARUBA_HOME}/zaruba" map set "${__ZRB_SINGLE_MAP_ENV}" "default" "${__ZRB_DEFAULT}")"
        __ZRB_MAP_ENVS="$("${ZARUBA_HOME}/zaruba" map set "${__ZRB_MAP_ENVS}" "${__ZRB_KEY}" "${__ZRB_SINGLE_MAP_ENV}")"
    done
    __ZRB_YAML_ENVS="$("${ZARUBA_HOME}/zaruba" yaml print "${__ZRB_MAP_ENVS}")"
    echo "${__ZRB_YAML_ENVS}"
}
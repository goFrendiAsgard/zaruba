_setReplacementMap() {
    _KEY="${1}"
    _VAL="${2}"
    _ZRB_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_REPLACEMENT_MAP}" "${_KEY}" "${_VAL}")"
}

_registerTask() {
    _PROJECT_FILE_NAME="${1}"
    _MAIN_TASK_NAME="${2}"
    _MODULE_TASK_NAME="${3}"
    if [ "$("${ZARUBA_HOME}/zaruba" task isExist "${_PROJECT_FILE_NAME}" "${_MODULE_TASK_NAME}")" = 1 ]
    then
        if [ "$("${ZARUBA_HOME}/zaruba" task isExist "${_PROJECT_FILE_NAME}" "${_MAIN_TASK_NAME}")" = 0 ]
        then
            "${ZARUBA_HOME}/zaruba" project addTask "${_PROJECT_FILE_NAME}" "${_MAIN_TASK_NAME}"
        fi
        "${ZARUBA_HOME}/zaruba" project addDependency "${_PROJECT_FILE_NAME}" "${_MAIN_TASK_NAME}" "[\"${_MODULE_TASK_NAME}\"]"
    fi
}
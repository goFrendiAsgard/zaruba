get_service_name_by_task_name() {
    _TASK=${1}
    _PATTERN="^run(.*)$"
    _TASK_SUBMATCH="$("${ZARUBA_HOME}/zaruba" str submatch "${_TASK}" "${_PATTERN}")"
    _SERVICE_PASCAL="$("${ZARUBA_HOME}/zaruba" list get "${_TASK_SUBMATCH}" 1)"
    _SERVICE="$("${ZARUBA_HOME}/zaruba" str toCamel "${_SERVICE_PASCAL}")"
    echo "${_SERVICE}"
}
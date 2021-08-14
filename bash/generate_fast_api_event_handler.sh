. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generate_fast_api_module.sh

generate_fast_api_event_handler() {
    _MODULE_TEMPLATE_LOCATION="${1}"
    _SERVICE_TEMPLATE_LOCATION="${2}"
    _TASK_TEMPLATE_LOCATION="${3}"
    _SERVICE_NAME="${4}"
    _MODULE_NAME="${5}"
    _EVENT_NAME="${6}"

    generate_fast_api_module \
        "${_MODULE_TEMPLATE_LOCATION}" \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}" \
        "${_MODULE_NAME}"
    
    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${SERVICE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${MODULE_NAME}")
    _SNAKE_EVENT_NAME=$("${ZARUBA_HOME}/zaruba" str toSnake "${EVENT_NAME}")
    _CAMEL_EVENT_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${EVENT_NAME}")

    # get controller lines
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" lines read "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" )
    _PATTERNS="$("${ZARUBA_HOME}/zaruba" list append "[]" \
        ".*def event_controller.*" \
    )"
    _LINE_INDEX=$("${ZARUBA_HOME}/zaruba" lines getIndex "${_CONTROLLER_LINES}" "${_PATTERNS}")

    # inject event handler
    _HANDLE_EVENT_PARTIAL=$(cat "${_MODULE_TEMPLATE_LOCATION}/partials/handle_event.py")
    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" map set "{}" \
        "zarubaEventName" "${_CAMEL_EVENT_NAME}" \
        "zaruba_event_name" "${_SNAKE_EVENT_NAME}" \
    )
    _HANDLE_EVENT_PARTIAL=$("${ZARUBA_HOME}/zaruba" str replace "${_HANDLE_EVENT_PARTIAL}" "${_REPLACEMENT_MAP}")
    _HANDLE_EVENT_PARTIAL="$("${ZARUBA_HOME}/zaruba" str indent "${_HANDLE_EVENT_PARTIAL}" "    ")"
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_CONTROLLER_LINES}" "${_LINE_INDEX}" "${_HANDLE_EVENT_PARTIAL}" )

    # save controller
    "${ZARUBA_HOME}/zaruba" lines write "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" "${_CONTROLLER_LINES}"
}


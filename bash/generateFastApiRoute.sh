. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generateFastApiModule.sh

generateFastApiRoute() {
    _MODULE_TEMPLATE_LOCATION="${1}"
    _SERVICE_TEMPLATE_LOCATION="${2}"
    _TASK_TEMPLATE_LOCATION="${3}"
    _SERVICE_NAME="${4}"
    _MODULE_NAME="${5}"
    _URL="${6}"
    _HTTP_METHOD="${6}"

    generateFastApiModule \
        "${_MODULE_TEMPLATE_LOCATION}" \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}" \
        "${_MODULE_NAME}"

    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${_SERVICE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str toCamel "${_MODULE_NAME}")
    _SNAKE_URL=$("${ZARUBA_HOME}/zaruba" str toSnake "${URL}")
    _LOWER_HTTP_METHOD=$("${ZARUBA_HOME}/zaruba" str toLower "${HTTP_METHOD}")

    # get controller lines
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" lines read "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" )
    _PATTERNS="$("${ZARUBA_HOME}/zaruba" list append "[]" \
        ".*def http_controller.*" \
    )"
    _LINE_INDEX=$("${ZARUBA_HOME}/zaruba" lines getIndex "${_CONTROLLER_LINES}" "${_PATTERNS}")

    # inject route handler
    _HANDLE_HTTP_PARTIAL=$(cat "${_MODULE_TEMPLATE_LOCATION}/partials/handle_http.py")
    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" map set "{}" \
        "zarubaUrl" "${URL}" \
        "zaruba_url" "${_SNAKE_URL}" \
        "zarubaHttpMethod" "${_LOWER_HTTP_METHOD}" \
    )
    _HANDLE_HTTP_PARTIAL=$("${ZARUBA_HOME}/zaruba" str replace "${_HANDLE_HTTP_PARTIAL}" "${_REPLACEMENT_MAP}")
    _HANDLE_HTTP_PARTIAL=$("${ZARUBA_HOME}/zaruba" str indent "${_HANDLE_HTTP_PARTIAL}" "    ")
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_CONTROLLER_LINES}" "${_LINE_INDEX}" "${_HANDLE_HTTP_PARTIAL}" )

    # save controller
    "${ZARUBA_HOME}/zaruba" lines write "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" "${_CONTROLLER_LINES}"
}


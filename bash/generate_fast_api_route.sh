. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generate_fast_api_module.sh

generate_fast_api_route() {
    _MODULE_TEMPLATE_LOCATION="${1}"
    _SERVICE_TEMPLATE_LOCATION="${2}"
    _TASK_TEMPLATE_LOCATION="${3}"
    _SERVICE_NAME="${4}"
    _MODULE_NAME="${5}"
    _URL="${6}"
    _HTTP_METHOD="${6}"

    generate_fast_api_module \
        "${_MODULE_TEMPLATE_LOCATION}" \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}" \
        "${_MODULE_NAME}"

    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${_SERVICE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${_MODULE_NAME}")
    _SNAKE_URL=$("${ZARUBA_HOME}/zaruba" strToSnake "${URL}")
    _LOWER_HTTP_METHOD=$("${ZARUBA_HOME}/zaruba" strToLower "${HTTP_METHOD}")

    # get controller lines
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" )
    _PATTERNS="$("${ZARUBA_HOME}/zaruba" appendToList "[]" \
        ".*def route_controller.*" \
    )"
    _LINE_INDEX=$("${ZARUBA_HOME}/zaruba" getLineIndex "${_CONTROLLER_LINES}" "${_PATTERNS}")

    # inject route handler
    _HANDLE_ROUTE_PARTIAL=$(cat "${_MODULE_TEMPLATE_LOCATION}/partials/handle_route.py")
    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zarubaUrl" "${URL}" \
        "zaruba_url" "${_SNAKE_URL}" \
        "zarubaHttpMethod" "${_LOWER_HTTP_METHOD}" \
    )
    _HANDLE_ROUTE_PARTIAL=$("${ZARUBA_HOME}/zaruba" strReplace "${_HANDLE_ROUTE_PARTIAL}" "${_REPLACEMENT_MAP}")
    _HANDLE_ROUTE_PARTIAL=$("${ZARUBA_HOME}/zaruba" strIndent "${_HANDLE_ROUTE_PARTIAL}" "    ")
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_CONTROLLER_LINES}" "${_LINE_INDEX}" "${_HANDLE_ROUTE_PARTIAL}" )

    # save controller
    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" "${_CONTROLLER_LINES}"
}


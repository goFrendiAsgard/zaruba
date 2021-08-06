. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generate_fast_api_module.sh

generate_fast_api_rpc_handler() {
    _MODULE_TEMPLATE_LOCATION="${1}"
    _SERVICE_TEMPLATE_LOCATION="${2}"
    _TASK_TEMPLATE_LOCATION="${3}"
    _SERVICE_NAME="${4}"
    _MODULE_NAME="${5}"
    _RPC_NAME="${6}"

    generate_fast_api_module \
        "${_MODULE_TEMPLATE_LOCATION}" \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}" \
        "${_MODULE_NAME}"
    
    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${SERVICE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${MODULE_NAME}")
    _SNAKE_RPC_NAME=$("${ZARUBA_HOME}/zaruba" strToSnake "${RPC_NAME}")
    _CAMEL_RPC_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${RPC_NAME}")

    # get controller lines
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" )
    _PATTERNS="$("${ZARUBA_HOME}/zaruba" appendToList "[]" \
        ".*def event_controller.*" \
    )"
    _LINE_INDEX=$("${ZARUBA_HOME}/zaruba" getLineIndex "${_CONTROLLER_LINES}" "${_PATTERNS}")

    # inject rpc handler
    _HANDLE_RPC_PARTIAL=$(cat "${_MODULE_TEMPLATE_LOCATION}/partials/handle_rpc.py")
    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zarubaRpcName" "${_CAMEL_RPC_NAME}" \
        "zaruba_rpc_name" "${_SNAKE_RPC_NAME}" \
    )
    _HANDLE_RPC_PARTIAL=$("${ZARUBA_HOME}/zaruba" strReplace "${_HANDLE_RPC_PARTIAL}" "${_REPLACEMENT_MAP}")
    _HANDLE_RPC_PARTIAL="$("${ZARUBA_HOME}/zaruba" strIndent "${_HANDLE_RPC_PARTIAL}" "    ")"
    _CONTROLLER_LINES=$("${ZARUBA_HOME}/zaruba" insertLineAfterIndex "${_CONTROLLER_LINES}" "${_LINE_INDEX}" "${_HANDLE_RPC_PARTIAL}" )

    # save controller
    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/${_CAMEL_MODULE_NAME}/controller.py" "${_CONTROLLER_LINES}"
}


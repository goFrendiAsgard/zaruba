_HANDLE_RPC_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiRpcHandler/partials/handle_rpc.py")"
_HANDLE_RPC_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_HANDLE_RPC_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/${_ZRB_APP_MODULE_NAME}/rpc.py"

_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CONTROLLER_FILE_LOCATION}")"

_PATTERN='["def register_rpc_handler"]'
_INSERT_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_LINES}" "${_PATTERN}")"

# get indentation
_INSERT_LINE="$("${ZARUBA_HOME}/zaruba" list get "${_LINES}" "${_INSERT_INDEX}")"
_INDENTATION="    $("${ZARUBA_HOME}/zaruba" str getIndentation "${_INSERT_LINE}")"
_INDENTED_HANDLE_RPC_SCRIPT="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_HANDLE_RPC_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_LINES}" "${_INSERT_INDEX}" "${_INDENTED_HANDLE_RPC_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_HOME}/zaruba" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"
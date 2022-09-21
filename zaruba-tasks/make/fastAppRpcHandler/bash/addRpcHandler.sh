echo "Registering rpc handler"

_HANDLE_RPC_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppRpcHandler/partials/handle_rpc.py")"
_HANDLE_RPC_SCRIPT="$("${ZARUBA_BIN}" str replace "${_HANDLE_RPC_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/modules/${_ZRB_APP_MODULE_NAME}/rpc.py"

_LINES="$("${ZARUBA_BIN}" lines read "${_CONTROLLER_FILE_LOCATION}")"

_PATTERN="def register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler"
_FUNCTION_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FUNCTION_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_FUNCTION_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_FUNCTION_INDEX}")"
_INDENTATION="    $("${ZARUBA_BIN}" str getIndentation "${_FUNCTION_LINE}")"
_INDENTED_HANDLE_RPC_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_HANDLE_RPC_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_FUNCTION_INDEX}" "${_INDENTED_HANDLE_RPC_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"

echo "Done registering rpc handler"
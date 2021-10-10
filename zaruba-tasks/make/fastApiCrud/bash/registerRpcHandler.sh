_IMPORT_RPC_HANDLER_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/import_rpc_handler.py")"
_IMPORT_RPC_HANDLER_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_RPC_HANDLER_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_HANDLE_RPC_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/handle_rpc.py")"
_HANDLE_RPC_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_HANDLE_RPC_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/${_ZRB_APP_MODULE_NAME}/rpc.py"

_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CONTROLLER_FILE_LOCATION}")"

# insert import
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertBefore "${_LINES}" 0 "${_IMPORT_RPC_HANDLER_SCRIPT}")"

# look for handler function
_FUNCTION_PATTERN='^(\s*)def register_rpc_handler\((.*)\)(.*)$'
_PATTERN="$("${ZARUBA_HOME}/zaruba" list append '[]' "${_FUNCTION_PATTERN}")"
_FUNCTION_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_LINES}" "${_PATTERN}")"
_FUNCTION_LINE="$("${ZARUBA_HOME}/zaruba" list get "${_LINES}" "${_FUNCTION_INDEX}")"
_FUNCTION_SUBMATCH="$("${ZARUBA_HOME}/zaruba" lines submatch "${_LINES}" "${_PATTERN}")"
_FUNCTION_INDENTATION="$("${ZARUBA_HOME}/zaruba" list get "${_FUNCTION_SUBMATCH}" 1)"
_FUNCTION_PARAM="$("${ZARUBA_HOME}/zaruba" list get "${_FUNCTION_SUBMATCH}" 2)"
_FUNCTION_SUFFIX="$("${ZARUBA_HOME}/zaruba" list get "${_FUNCTION_SUBMATCH}" 3)"
_NEW_FUNCTION_LINE="${_FUNCTION_INDENTATION}def register_rpc_handler(${_FUNCTION_PARAM}, ${_ZRB_APP_CRUD_SNAKE_ENTITY}_repo: ${_ZRB_APP_CRUD_PASCAL_ENTITY}Repo)${_FUNCTION_SUFFIX}"

# replace function signature
_LINES="$("${ZARUBA_HOME}/zaruba" list set "${_LINES}" "${_FUNCTION_INDEX}" "${_NEW_FUNCTION_LINE}")"

# get indentation
_INDENTATION="    ${_FUNCTION_INDENTATION}"
_INDENTED_HANDLE_RPC_SCRIPT="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_HANDLE_RPC_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_LINES}" "${_FUNCTION_INDEX}" "${_INDENTED_HANDLE_RPC_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_HOME}/zaruba" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"
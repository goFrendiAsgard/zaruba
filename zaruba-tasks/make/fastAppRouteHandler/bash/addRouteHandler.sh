echo "Registering route handler"

_HANDLE_ROUTE_SCRIPT="$(_readText "${ZARUBA_HOME}/zaruba-tasks/make/fastAppRouteHandler/partials/handle_route.py")"
_HANDLE_ROUTE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_HANDLE_ROUTE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/route.py"

_LINES="$(_readLines "${_CONTROLLER_FILE_LOCATION}")"

_PATTERN="def register_${_ZRB_SNAKE_APP_MODULE_NAME}_api_route"
_FUNCTION_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FUNCTION_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_FUNCTION_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_FUNCTION_INDEX}")"
_INDENTATION="    $("${ZARUBA_BIN}" str getIndentation "${_FUNCTION_LINE}")"
_INDENTED_HANDLE_ROUTE_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_HANDLE_ROUTE_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_INDENTED_HANDLE_ROUTE_SCRIPT}" --index="${_FUNCTION_INDEX}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_LINES}" "${_CONTROLLER_FILE_LOCATION}"

echo "Done registering route handler"
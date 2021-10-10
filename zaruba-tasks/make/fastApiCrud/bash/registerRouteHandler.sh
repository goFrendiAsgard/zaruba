_IMPORT_ROUTE_HANDLER_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/import_route_handler.py")"
_IMPORT_ROUTE_HANDLER_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_ROUTE_HANDLER_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_HANDLE_ROUTE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/handle_route.py")"
_HANDLE_ROUTE_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_HANDLE_ROUTE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/${_ZRB_APP_MODULE_NAME}/route.py"

_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CONTROLLER_FILE_LOCATION}")"

# insert import
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertBefore "${_LINES}" 0 "${_IMPORT_ROUTE_HANDLER_SCRIPT}")"

# look for handler function
_PATTERN='["def register_route_handler"]'
_FUNCTION_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_LINES}" "${_PATTERN}")"

# get indentation
_FUNCTION_LINE="$("${ZARUBA_HOME}/zaruba" list get "${_LINES}" "${_FUNCTION_INDEX}")"
_INDENTATION="    $("${ZARUBA_HOME}/zaruba" str getIndentation "${_FUNCTION_LINE}")"
_INDENTED_HANDLE_ROUTE_SCRIPT="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_HANDLE_ROUTE_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_LINES}" "${_FUNCTION_INDEX}" "${_INDENTED_HANDLE_ROUTE_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_HOME}/zaruba" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"
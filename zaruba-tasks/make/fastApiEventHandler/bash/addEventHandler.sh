_HANDLE_EVENT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiEventHandler/partials/handle_event.py")"
_HANDLE_EVENT_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_HANDLE_EVENT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/${_ZRB_APP_MODULE_NAME}/event.py"

_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_CONTROLLER_FILE_LOCATION}")"

_PATTERN='["def register_event_handler"]'
_FUNCTION_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_LINES}" "${_PATTERN}")"

# get indentation
_FUNCTION_LINE="$("${ZARUBA_HOME}/zaruba" list get "${_LINES}" "${_FUNCTION_INDEX}")"
_INDENTATION="    $("${ZARUBA_HOME}/zaruba" str getIndentation "${_FUNCTION_LINE}")"
_INDENTED_HANDLE_EVENT_SCRIPT="$("${ZARUBA_HOME}/zaruba" str fullIndent "${_HANDLE_EVENT_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertAfter "${_LINES}" "${_FUNCTION_INDEX}" "${_INDENTED_HANDLE_EVENT_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_HOME}/zaruba" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"
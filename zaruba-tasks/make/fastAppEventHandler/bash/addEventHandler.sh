_HANDLE_EVENT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppEventHandler/partials/handle_event.py")"
_HANDLE_EVENT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_HANDLE_EVENT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/${_ZRB_APP_MODULE_NAME}/event.py"

_LINES="$("${ZARUBA_BIN}" lines read "${_CONTROLLER_FILE_LOCATION}")"

_PATTERN="def register_${_ZRB_SNAKE_APP_MODULE_NAME}_event_handler"
_FUNCTION_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FUNCTION_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_FUNCTION_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_FUNCTION_INDEX}")"
_INDENTATION="    $("${ZARUBA_BIN}" str getIndentation "${_FUNCTION_LINE}")"
_INDENTED_HANDLE_EVENT_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_HANDLE_EVENT_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_FUNCTION_INDEX}" "${_INDENTED_HANDLE_EVENT_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"

echo "Done adding event handler"

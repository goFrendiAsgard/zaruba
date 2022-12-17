echo "Adding event handler"

_addEventHandler() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/event.py"
    _PATTERN="def register_${_ZRB_SNAKE_APP_MODULE_NAME}_event_handler"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"
    _INDENTATION="$(_getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppEventHandler/partials/handle_event.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "    ${_INDENTATION}")"

    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_DESTINATION}"
}

_addEventHandler

echo "Done adding event handler"

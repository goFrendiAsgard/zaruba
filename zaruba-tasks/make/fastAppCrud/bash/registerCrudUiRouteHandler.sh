set -e
echo "Registering CRUD UI route handler"

# module/<module_name>/route.py
_registerCrudUiRouteHandler() {
    _FILE_PATH="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/route.py"
    _PATTERN="menu_service.add_menu\("
    _LINE_INDEX="$(_getLineIndexFromFile "${_FILE_PATH}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_FILE_PATH}" "${_LINE_INDEX}")"
    _INDENTATION="$(_getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/handle_ui_route.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "${_INDENTATION}")"

    _insertPartialAfter "${_FILE_PATH}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_FILE_PATH}"
}

_registerCrudUiRouteHandler

echo "Done registering CRUD UI route handler"
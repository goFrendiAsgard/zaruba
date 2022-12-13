echo "Registering page route handler"

_HANDLE_ROUTE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppPage/partials/handle_route.py")"
_HANDLE_ROUTE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_HANDLE_ROUTE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/route.py"

_LINES="$("${ZARUBA_BIN}" lines read "${_CONTROLLER_FILE_LOCATION}")"

_PATTERN="menu_service.add_menu\("
_ADD_MENU_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_ADD_MENU_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_ADD_MENU_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_ADD_MENU_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_ADD_MENU_LINE}")"
_INDENTED_HANDLE_ROUTE_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_HANDLE_ROUTE_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_INDENTED_HANDLE_ROUTE_SCRIPT}" --index="${_ADD_MENU_INDEX}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_LINES}" "${_CONTROLLER_FILE_LOCATION}"

echo "Done registering page route handler"
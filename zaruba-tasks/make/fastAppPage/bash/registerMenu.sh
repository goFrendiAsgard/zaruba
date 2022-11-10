echo "Registering menu"

_REGISTER_MENU_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppPage/partials/register_menu.py")"
_REGISTER_MENU_SCRIPT="$("${ZARUBA_BIN}" str replace "${_REGISTER_MENU_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CREATE_MENU_SERVICE_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/configs/menuServiceFactory.py"

_LINES="$("${ZARUBA_BIN}" lines read "${_CREATE_MENU_SERVICE_FILE_LOCATION}")"

_PATTERN="menu_service\.add_menu\(name=\'${_ZRB_SNAKE_APP_MODULE_NAME}\',"
_PARENT_MENU_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_PARENT_MENU_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_PARENT_MENU_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_PARENT_MENU_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_PARENT_MENU_LINE}")"
_INDENTED_REGISTER_MENU_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_REGISTER_MENU_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_PARENT_MENU_INDEX}" "${_INDENTED_REGISTER_MENU_SCRIPT}")"

chmod 755 "${_CREATE_MENU_SERVICE_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_CREATE_MENU_SERVICE_FILE_LOCATION}" "${_LINES}"

echo "Done registering menu"
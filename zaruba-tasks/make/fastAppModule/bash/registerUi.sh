echo "Registering user interface"

_REGISTER_UI_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/register_ui.py")"
_REGISTER_UI_SCRIPT="$("${ZARUBA_BIN}" str replace "${_REGISTER_UI_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CREATE_MENU_SERVICE_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/configs/menuServiceFactory.py"

_LINES="$("${ZARUBA_BIN}" lines read "${_CREATE_MENU_SERVICE_FILE_LOCATION}")"

_PATTERN="return menu_service"
_RETURN_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_RETURN_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_RETURN_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_RETURN_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_RETURN_LINE}")"
_INDENTED_REGISTER_UI_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_REGISTER_UI_SCRIPT}" "${_INDENTATION}")"

# insert handler
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_RETURN_INDEX}" "${_INDENTED_REGISTER_UI_SCRIPT}")"

chmod 755 "${_CREATE_MENU_SERVICE_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_CREATE_MENU_SERVICE_FILE_LOCATION}" "${_LINES}"

echo "Done registering user interface"
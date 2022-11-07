set -e
echo "Registering repo and service"

_IMPORT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_repo_and_service.py")"
_IMPORT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_IMPORT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_INIT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/init_repo_and_service.py")"
_INIT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_INIT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"


####################################################################
## main.py
####################################################################

_MAIN_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/main.py"
_MAIN_LINES="$("${ZARUBA_BIN}" lines read "${_MAIN_FILE_LOCATION}")"

####################################################################
# insert import

_MAIN_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_MAIN_LINES}" 0 "${_IMPORT_SCRIPT}")"

# init repo
_PATTERN="if enable_${_ZRB_SNAKE_APP_MODULE_NAME}_module:"
_ENGINE_DECLARATION_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_MAIN_LINES}" "${_PATTERN}")"
if [ "${_ENGINE_DECLARATION_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_INIT_SCRIPT_INDENTATION="    "
_INDENTED_INIT_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_INIT_SCRIPT}" "${_INIT_SCRIPT_INDENTATION}")"
_MAIN_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_MAIN_LINES}" "${_ENGINE_DECLARATION_INDEX}" "${_INDENTED_INIT_SCRIPT}")"

####################################################################
# look for rpc call

_PATTERN="^(\s*)register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler\((.*)\)(.*)$"
_CALL_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_MAIN_LINES}" "${_PATTERN}")"
if [ "${_CALL_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_CALL_LINE="$("${ZARUBA_BIN}" list get "${_MAIN_LINES}" "${_CALL_INDEX}")"
_CALL_SUBMATCH="$("${ZARUBA_BIN}" lines submatch "${_MAIN_LINES}" "${_PATTERN}")"
_CALL_INDENTATION="$("${ZARUBA_BIN}" list get "${_CALL_SUBMATCH}" 1)"
_CALL_PARAM="$("${ZARUBA_BIN}" list get "${_CALL_SUBMATCH}" 2)"
_CALL_SUFFIX="$("${ZARUBA_BIN}" list get "${_CALL_SUBMATCH}" 3)"
_NEW_CALL_LINE="${_CALL_INDENTATION}register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler(${_CALL_PARAM}, ${_ZRB_SNAKE_APP_CRUD_ENTITY}_service)${_CALL_SUFFIX}"

# replace rpc call
_MAIN_LINES="$("${ZARUBA_BIN}" list set "${_MAIN_LINES}" "${_CALL_INDEX}" "${_NEW_CALL_LINE}")"

chmod 755 "${_MAIN_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_MAIN_FILE_LOCATION}" "${_MAIN_LINES}"

echo "Done registering repo and service"
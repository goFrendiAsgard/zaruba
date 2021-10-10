set -x
_IMPORT_REPO_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/import_repo.py")"
_IMPORT_REPO_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_REPO_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_MAIN_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/main.py"
_LINES="$("${ZARUBA_HOME}/zaruba" lines read "${_MAIN_FILE_LOCATION}")"

# insert import
_LINES="$("${ZARUBA_HOME}/zaruba" lines insertBefore "${_LINES}" 0 "${_IMPORT_REPO_SCRIPT}")"

# look for rpc call
_CALL_PATTERN="^(\s*)register_${_ZRB_APP_SNAKE_MODULE_NAME}_rpc_handler\((.*)\)(.*)$"
_PATTERN="$("${ZARUBA_HOME}/zaruba" list append '[]' "${_CALL_PATTERN}")"
_CALL_INDEX="$("${ZARUBA_HOME}/zaruba" lines getIndex "${_LINES}" "${_PATTERN}")"
_CALL_LINE="$("${ZARUBA_HOME}/zaruba" list get "${_LINES}" "${_CALL_INDEX}")"
_CALL_SUBMATCH="$("${ZARUBA_HOME}/zaruba" lines submatch "${_LINES}" "${_PATTERN}")"
_CALL_INDENTATION="$("${ZARUBA_HOME}/zaruba" list get "${_CALL_SUBMATCH}" 1)"
_CALL_PARAM="$("${ZARUBA_HOME}/zaruba" list get "${_CALL_SUBMATCH}" 2)"
_CALL_SUFFIX="$("${ZARUBA_HOME}/zaruba" list get "${_CALL_SUBMATCH}" 3)"
_NEW_CALL_LINE="${_CALL_INDENTATION}register_${_ZRB_APP_SNAKE_MODULE_NAME}_rpc_handler(${_CALL_PARAM}, ${_ZRB_APP_CRUD_SNAKE_ENTITY}_repo)${_CALL_SUFFIX}"

# replace call signature
_LINES="$("${ZARUBA_HOME}/zaruba" list set "${_LINES}" "${_CALL_INDEX}" "${_NEW_CALL_LINE}")"

chmod 755 "${_MAIN_FILE_LOCATION}"
"${ZARUBA_HOME}/zaruba" lines write "${_MAIN_FILE_LOCATION}" "${_LINES}"

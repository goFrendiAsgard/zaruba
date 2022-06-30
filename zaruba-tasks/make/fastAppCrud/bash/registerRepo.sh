set -e
echo "Registering repo"

_IMPORT_REPO_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_repo.py")"
_IMPORT_REPO_SCRIPT="$("${ZARUBA_BIN}" str replace "${_IMPORT_REPO_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_INIT_REPO_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/init_repo.py")"
_INIT_REPO_SCRIPT="$("${ZARUBA_BIN}" str replace "${_INIT_REPO_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"


####################################################################
####################################################################
## Updating main.py
####################################################################
####################################################################

_MAIN_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/main.py"
_MAIN_LINES="$("${ZARUBA_BIN}" lines read "${_MAIN_FILE_LOCATION}")"

####################################################################
# insert import
####################################################################
_MAIN_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_MAIN_LINES}" 0 "${_IMPORT_REPO_SCRIPT}")"

# init repo
_ENGINE_DECLARATION_PATTERN="engine(\s*)="
_PATTERN="$("${ZARUBA_BIN}" list append '[]' "${_ENGINE_DECLARATION_PATTERN}")"
_ENGINE_DECLARATION_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_MAIN_LINES}" "${_PATTERN}")"
if [ "${_ENGINE_DECLARATION_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_MAIN_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_MAIN_LINES}" "${_ENGINE_DECLARATION_INDEX}" "${_INIT_REPO_SCRIPT}")"

####################################################################
# look for rpc call
####################################################################
_CALL_PATTERN="^(\s*)register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler\((.*)\)(.*)$"
_PATTERN="$("${ZARUBA_BIN}" list append '[]' "${_CALL_PATTERN}")"
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
_NEW_CALL_LINE="${_CALL_INDENTATION}register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler(${_CALL_PARAM}, ${_ZRB_SNAKE_APP_CRUD_ENTITY}_repo)${_CALL_SUFFIX}"

# replace rpc call
_MAIN_LINES="$("${ZARUBA_BIN}" list set "${_MAIN_LINES}" "${_CALL_INDEX}" "${_NEW_CALL_LINE}")"

chmod 755 "${_MAIN_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_MAIN_FILE_LOCATION}" "${_MAIN_LINES}"


####################################################################
####################################################################
## Updating alembic/env.py
####################################################################
####################################################################

_ENV_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/alembic/env.py"
_ENV_LINES="$("${ZARUBA_BIN}" lines read "${_ENV_FILE_LOCATION}")"

####################################################################
# insert import
####################################################################
_ENV_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_ENV_LINES}" 1 "${_IMPORT_REPO_SCRIPT}")"
chmod 755 "${_ENV_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_ENV_FILE_LOCATION}" "${_ENV_LINES}"

echo "Done registering repo"
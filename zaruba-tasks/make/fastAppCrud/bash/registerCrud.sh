set -e
echo "Registering CRUD"

_importRepoAndService() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/main.py"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_repo_and_service.py")"
    _insertPartialBefore "${_DESTINATION}" "${_NEW_CONTENT}" 0
    chmod 755 "${_DESTINATION}"
}

_registerRepoAndService() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/main.py"
    _PATTERN="if enable_${_ZRB_SNAKE_APP_MODULE_NAME}_module:"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"
    _INDENTATION="$(_getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/init_repo_and_service.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "    ${_INDENTATION}")"

    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_DESTINATION}"
}

_updateRpcCall() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/main.py"
    _PATTERN="^(\s*)register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler\((.*)\)(.*)$"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

    _SUBMATCH="$("${ZARUBA_BIN}" str submatch "${_LINE}" "${_PATTERN}")"

    _INDENTATION="$("${ZARUBA_BIN}" list get "${_SUBMATCH}" 1)"
    _PARAM="$("${ZARUBA_BIN}" list get "${_SUBMATCH}" 2)"
    _SUFFIX="$("${ZARUBA_BIN}" list get "${_SUBMATCH}" 3)"

    _NEW_LINE="${_INDENTATION}register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler(${_PARAM}, ${_ZRB_SNAKE_APP_CRUD_ENTITY}_service)${_SUFFIX}"
    _replacePartialAt "${_DESTINATION}" "${_NEW_LINE}" "${_LINE_INDEX}"

    chmod 755 "${_DESTINATION}"
}

echo "Import repo and service"
_importRepoAndService
echo "Register repo and service"
_registerRepoAndService
echo "Update RPC call"
_updateRpcCall

echo "Done registering CRUD"
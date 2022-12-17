set -e
echo "Registering CRUD RPC handler"

_importCrudRpcHandler() {
    _FILE_PATH="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/rpc.py"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_rpc_handler.py")"
    _insertPartialBefore "${_FILE_PATH}" "${_NEW_CONTENT}" 0
    chmod 755 "${_FILE_PATH}"
}

_registerCrudRpcHandler() {
    _FILE_PATH="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/rpc.py"
    _PATTERN="def register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler"
    _LINE_INDEX="$(_getLineIndexFromFile "${_FILE_PATH}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_FILE_PATH}" "${_LINE_INDEX}")"
    _INDENTATION="$(_getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/handle_rpc.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "    ${_INDENTATION}")"

    _insertPartialAfter "${_FILE_PATH}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_FILE_PATH}"
}

_updateCrudRpcCall() {
    _FILE_PATH="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/rpc.py"
    _PATTERN="^(\s*)def register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler\((.*)\)(.*)$"
    _LINE_INDEX="$(_getLineIndexFromFile "${_FILE_PATH}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_FILE_PATH}" "${_LINE_INDEX}")"

    _SUBMATCH="$("${ZARUBA_BIN}" str submatch "${_LINE}" "${_PATTERN}")"

    _INDENTATION="$("${ZARUBA_BIN}" list get "${_SUBMATCH}" 1)"
    _PARAM="$("${ZARUBA_BIN}" list get "${_SUBMATCH}" 2)"
    _SUFFIX="$("${ZARUBA_BIN}" list get "${_SUBMATCH}" 3)"

    _NEW_LINE="${_INDENTATION}def register_${_ZRB_SNAKE_APP_MODULE_NAME}_rpc_handler(${_PARAM}, ${_ZRB_SNAKE_APP_CRUD_ENTITY}_service: ${_ZRB_PASCAL_APP_CRUD_ENTITY}Service)${_SUFFIX}"
    _replacePartialAt "${_FILE_PATH}" "${_NEW_LINE}" "${_LINE_INDEX}"

    chmod 755 "${_FILE_PATH}"
}

echo "Import CRUD RPC handler"
_importCrudRpcHandler
echo "Register CRUD RPC handler"
_registerCrudRpcHandler
echo "Update CRUD RPC call"
_updateCrudRpcCall

echo "Done registering CRUD RPC handler"
set -e
echo "Updating test field"


_addTestFieldDeclaration() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/${_ZRB_SNAKE_APP_CRUD_ENTITY}/test_${_ZRB_SNAKE_APP_CRUD_ENTITY}_service_util.py"
    _PATTERN="[\t ]*(dummy_${_ZRB_SNAKE_APP_CRUD_ENTITY}_data[\t ]*=[\t ]*${_ZRB_PASCAL_APP_CRUD_ENTITY}Data\([\t ]*)"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}")"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

    _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/test_field.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "    ${_INDENTATION}")"

    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_DESTINATION}"
}

_addTestFieldDeclaration

echo "Done updating test field"
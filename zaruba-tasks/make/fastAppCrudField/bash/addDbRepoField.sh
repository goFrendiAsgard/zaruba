set -e
echo "Updating db repo"

_addFieldDeclaration() {
    _DESTINATION="${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/${_ZRB_APP_CRUD_ENTITY}/repo/db_${_ZRB_SNAKE_APP_CRUD_ENTITY}_repo.py"
    _PATTERN="$("${ZARUBA_BIN}" list append '[]' "class[\t ]+DB${_ZRB_PASCAL_APP_CRUD_ENTITY}Entity\(")"
    _PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "__tablename__[\t ]*=")"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

    _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/repo_field_declaration.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "${_INDENTATION}")"

    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_DESTINATION}"
}


echo "Add field declaration"
_addFieldDeclaration

echo "Done updating db repo"
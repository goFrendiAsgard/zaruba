set -e
echo "Adding schema field declaration"

_addSchemaFieldDeclaration() {
    _DESTINATION="${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/schema/${_ZRB_SNAKE_APP_CRUD_ENTITY}.py"
    _PATTERN="$("${ZARUBA_BIN}" list append '[]' "class[\t ]*${_ZRB_PASCAL_APP_CRUD_ENTITY}Data\(")"
    _PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "created_at[\t ]*:[\t ]Optional")"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        exit 1
    fi
    _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

    _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_LINE}")"

    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/schema_field_declaration.py")"
    _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "${_INDENTATION}")"

    _insertPartialBefore "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
    chmod 755 "${_DESTINATION}"
}

_addSchemaFieldDeclaration

echo "Done adding schema field declaration"
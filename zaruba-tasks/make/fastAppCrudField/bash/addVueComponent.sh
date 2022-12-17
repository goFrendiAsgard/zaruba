echo "Updating vue component"

_addFormInput() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/public/vue/modules/${_ZRB_APP_MODULE_NAME}/crud/${_ZRB_APP_CRUD_ENTITIES}Crud.vue"
    _PATTERN="<!-- Put form input here"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        _PATTERN="$("${ZARUBA_BIN}" list append '[]' "[ \t]*<div.*[ \t]+class[ \t]*=[ \t]*\".*modal-body.*\"*")"
        _PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "[ \t]*<div.*[ \t]+class[ \t]*=[ \t]*\".*mb-3.*\"*")"
        _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
        if [ "${_LINE_INDEX}" = "-1" ]
        then
            echo "Pattern not found: ${_PATTERN}"
            echo "${_RED}${_BOLD}Cannot add form input${_NORMAL}"
        fi
    fi
    if [ "${_LINE_INDEX}" != "-1" ]
    then
        echo "add form input"
        _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

        _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_LINE}")"

        _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/vue/form_input.html")"
        _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "${_INDENTATION}")"

        _insertPartialBefore "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
        chmod 755 "${_DESTINATION}"
    fi
}

_addColumnHeader() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/public/vue/modules/${_ZRB_APP_MODULE_NAME}/crud/${_ZRB_APP_CRUD_ENTITIES}Crud.vue"
    _PATTERN="<!-- Put column header here"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        _PATTERN="[ \t]*<th.*[ \t]+id[ \t]*=[ \t]*\"th-action\".*"
        _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
        if [ "${_LINE_INDEX}" = "-1" ]
        then
            echo "Pattern not found: ${_PATTERN}"
            echo "${_RED}${_BOLD}Cannot add column header${_NORMAL}"
        fi
    fi
    if [ "${_LINE_INDEX}" != "-1" ]
    then
        echo "add column header"
        _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

        _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_LINE}")"

        _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/vue/column_header.html")"
        _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "${_INDENTATION}")"

        _insertPartialBefore "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
        chmod 755 "${_DESTINATION}"
    fi
}

_addColumnValue() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/public/vue/modules/${_ZRB_APP_MODULE_NAME}/crud/${_ZRB_APP_CRUD_ENTITIES}Crud.vue"
    _PATTERN="<!-- Put column value here"
    _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
    if [ "${_LINE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        _PATTERN="[ \t]*<th.*[ \t]+id[ \t]*=[ \t]*\"th-action\".*"
        _LINE_INDEX="$(_getLineIndexFromFile "${_DESTINATION}" "${_PATTERN}" --index=-1)"
        if [ "${_LINE_INDEX}" = "-1" ]
        then
            echo "Pattern not found: ${_PATTERN}"
            echo "${_RED}${_BOLD}Cannot add column value${_NORMAL}"
        fi
    fi
    if [ "${_LINE_INDEX}" != "-1" ]
    then
        echo "add column value"
        _LINE="$(_getLineFromFile "${_DESTINATION}" "${_LINE_INDEX}")"

        _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_LINE}")"

        _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/vue/column_value.html")"
        _NEW_CONTENT="$(_indent "${_NEW_CONTENT}" "${_INDENTATION}")"

        _insertPartialBefore "${_DESTINATION}" "${_NEW_CONTENT}" "${_LINE_INDEX}"
        chmod 755 "${_DESTINATION}"
    fi
}

echo "Add form input"
_addFormInput
echo "Add column header"
_addColumnHeader
echo "Add column value"
_addColumnValue

echo "Done updating vue component"
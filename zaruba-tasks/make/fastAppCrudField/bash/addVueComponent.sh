echo "Updating vue component"

_FORM_INPUT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/vue/form_input.html")"
_FORM_INPUT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FORM_INPUT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_COLUMN_HEADER_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/vue/column_header.html")"
_COLUMN_HEADER_SCRIPT="$("${ZARUBA_BIN}" str replace "${_COLUMN_HEADER_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_COLUMN_VALUE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/vue/column_value.html")"
_COLUMN_VALUE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_COLUMN_VALUE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

#########################################################
# Read existing vue component

_UI_COMPONENT_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/public/vue/modules/${_ZRB_APP_MODULE_NAME}/crud/${_ZRB_APP_CRUD_ENTITIES}Crud.vue"

_LINES="$("${ZARUBA_BIN}" lines read "${_UI_COMPONENT_FILE_LOCATION}")"

#########################################################
# FORM_INPUT

_IS_PATTERN_FOUND=1
_PATTERN="<!-- Put form input here"
_FORM_INPUT_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FORM_INPUT_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    _PATTERN="$("${ZARUBA_BIN}" list append '[]' "[ \t]*<div.*[ \t]+class[ \t]*=[ \t]*\".*modal-body.*\"*")"
    _PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "[ \t]*<div.*[ \t]+class[ \t]*=[ \t]*\".*mb-3.*\"*")"
    _FORM_INPUT_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
    if [ "${_FORM_INPUT_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        echo "${_RED}${_BOLD}Cannot create form input${_NORMAL}"
        _IS_PATTERN_FOUND=0
    fi
fi

if [ "${_IS_PATTERN_FOUND}" = 1 ]
then
    echo "Inject form input"
    # get indentation
    _FORM_INPUT_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_FORM_INPUT_INDEX}")"
    _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_FORM_INPUT_LINE}")"
    _INDENTED_FORM_INPUT_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_FORM_INPUT_SCRIPT}" "${_INDENTATION}")"
    _LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_INDENTED_FORM_INPUT_SCRIPT}" --index="${_FORM_INPUT_INDEX}")"
fi


#########################################################
# COLUMN HEADER

_IS_PATTERN_FOUND=1
_PATTERN="<!-- Put column header here"
_COLUMN_HEADER_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_COLUMN_HEADER_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    _PATTERN="[ \t]*<th.*[ \t]+id[ \t]*=[ \t]*\"th-action\".*"
    _COLUMN_HEADER_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
    if [ "${_COLUMN_HEADER_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        echo "${_RED}${_BOLD}Cannot create column header${_NORMAL}"
        _IS_PATTERN_FOUND=0
    fi
fi

if [ "${_IS_PATTERN_FOUND}" = 1 ]
then
    echo "Inject column header"
    # get indentation
    _COLUMN_HEADER_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_COLUMN_HEADER_INDEX}")"
    _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_COLUMN_HEADER_LINE}")"
    _INDENTED_COLUMN_HEADER_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_COLUMN_HEADER_SCRIPT}" "${_INDENTATION}")"
    _LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_INDENTED_COLUMN_HEADER_SCRIPT}" --index="${_COLUMN_HEADER_INDEX}")"
fi


#########################################################
# COLUMN VALUE

_IS_PATTERN_FOUND=1
_PATTERN="<!-- Put column value here"
_COLUMN_VALUE_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_COLUMN_VALUE_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    _PATTERN="[ \t]*<td.*[ \t]+id[ \t]*=[ \t]*\"td-action\".*"
    _COLUMN_VALUE_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
    if [ "${_COLUMN_VALUE_INDEX}" = "-1" ]
    then
        echo "Pattern not found: ${_PATTERN}"
        echo "${_RED}${_BOLD}Cannot create column value${_NORMAL}"
        _IS_PATTERN_FOUND=0
    fi
fi

if [ "${_IS_PATTERN_FOUND}" = 1 ]
then
    echo "Inject column value"
    # get indentation
    _COLUMN_VALUE_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_COLUMN_VALUE_INDEX}")"
    _INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_COLUMN_VALUE_LINE}")"
    _INDENTED_COLUMN_VALUE_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_COLUMN_VALUE_SCRIPT}" "${_INDENTATION}")"
    _LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_INDENTED_COLUMN_VALUE_SCRIPT}" --index="${_COLUMN_VALUE_INDEX}")"
fi

#########################################################
# Overwrite existing vue component

chmod 755 "${_UI_COMPONENT_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_LINES}" "${_UI_COMPONENT_FILE_LOCATION}"

echo "Done updating vue component"
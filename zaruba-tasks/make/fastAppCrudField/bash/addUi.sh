echo "Updating jinja template"

_FORM_INPUT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/jinja/form_input.html")"
_FORM_INPUT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FORM_INPUT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_TABLE_HEADER_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/jinja/table_header.html")"
_TABLE_HEADER_SCRIPT="$("${ZARUBA_BIN}" str replace "${_TABLE_HEADER_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_TABLE_VALUE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/jinja/table_value.html")"
_TABLE_VALUE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_TABLE_VALUE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

#########################################################
# Read existing jinja template

_JINJA_TEMPLATE_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/_jinja_templates/${_ZRB_APP_MODULE_NAME}/crud/${_ZRB_APP_CRUD_ENTITIES}.html"

_LINES="$("${ZARUBA_BIN}" lines read "${_JINJA_TEMPLATE_FILE_LOCATION}")"

#########################################################
# FORM_INPUT

_PATTERN="$("${ZARUBA_BIN}" list append '[]' "[ \t]*<div.*[ \t]+class[ \t]*=[ \t]*\".*modal-body.*\"*")"
_PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "[ \t]*<div.*[ \t]+class[ \t]*=[ \t]*\".*mb-3.*\"*")"
_FORM_INPUT_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FORM_INPUT_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_FORM_INPUT_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_FORM_INPUT_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_FORM_INPUT_LINE}")"
_INDENTED_FORM_INPUT_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_FORM_INPUT_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_FORM_INPUT_INDEX}" "${_INDENTED_FORM_INPUT_SCRIPT}")"


#########################################################
# TH

_PATTERN="[ \t]*<th.*[ \t]+id[ \t]*=[ \t]*\"th-action\".*"
_TH_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_TH_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_TH_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_TH_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_TH_LINE}")"
_INDENTED_TABLE_HEADER_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_TABLE_HEADER_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_TH_INDEX}" "${_INDENTED_TABLE_HEADER_SCRIPT}")"


#########################################################
# TD

_PATTERN="[ \t]*<td.*[ \t]+id[ \t]*=[ \t]*\"td-action\".*"
_TD_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_TD_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

# get indentation
_TD_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_TD_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_TD_LINE}")"
_INDENTED_TABLE_VALUE_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_TABLE_VALUE_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_TD_INDEX}" "${_INDENTED_TABLE_VALUE_SCRIPT}")"

#########################################################
# Overwrite existing jinja template

chmod 755 "${_JINJA_TEMPLATE_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_JINJA_TEMPLATE_FILE_LOCATION}" "${_LINES}"

echo "Done updating jinja template"
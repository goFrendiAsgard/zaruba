set -e
echo "Updating schema field declaration"

_FIELD_SCRIPT_TEMPLATE="$(_readText "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/schema_field_declaration.py")"
_FIELD_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FIELD_SCRIPT_TEMPLATE}" "${_ZRB_REPLACEMENT_MAP}")"

#########################################################
# Read existing schema

_SCHEMA_LOCATION="${_ZRB_APP_DIRECTORY}/schema/${_ZRB_SNAKE_APP_CRUD_ENTITY}.py"
_LINES="$(_readLines "${_SCHEMA_LOCATION}")"


_PATTERN="$("${ZARUBA_BIN}" list append '[]' "class[\t ]*${_ZRB_PASCAL_APP_CRUD_ENTITY}Data\(")"
_PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "created_at[\t ]*:[\t ]Optional")"
_SCHEMA_CLASS_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_SCHEMA_CLASS_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_SCHEMA_CLASS_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_SCHEMA_CLASS_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_SCHEMA_CLASS_LINE}")"
_INDENTED_FIELD_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_FIELD_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_INDENTED_FIELD_SCRIPT}" --index="${_SCHEMA_CLASS_INDEX}")"

#########################################################
# Overwrite existing schema

chmod 755 "${_SCHEMA_LOCATION}"
"${ZARUBA_BIN}" lines write "${_LINES}" "${_SCHEMA_LOCATION}"

echo "Done updating schema field declaration"
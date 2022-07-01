set -e
echo "Updating mem repo"

_FIELD_INSERT_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/repo_field_insert.py")"
_FIELD_INSERT_SCRIPT="    $("${ZARUBA_BIN}" str replace "${_FIELD_INSERT_SCRIPT_TEMPLATE}" "${_ZRB_REPLACEMENT_MAP}")"

_FIELD_UPDATE_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/repo_field_update_mem.py")"
_FIELD_UPDATE_SCRIPT="    $("${ZARUBA_BIN}" str replace "${_FIELD_UPDATE_SCRIPT_TEMPLATE}" "${_ZRB_REPLACEMENT_MAP}")"

#########################################################
# Read existing repo

_REPO_LOCATION="${_ZRB_APP_DIRECTORY}/repos/${_ZRB_APP_CRUD_ENTITY}.py"
_LINES="$("${ZARUBA_BIN}" lines read "${_REPO_LOCATION}")"

#########################################################
# Insert

_PATTERN="$("${ZARUBA_BIN}" list append '[]' "class[\t ]*Mem${_ZRB_PASCAL_APP_CRUD_ENTITY}Repo\(")"
_PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "def[ \t]+insert\(")"
_PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "created_at[\t ]*=[\t ]*datetime\.datetime\.utcnow")"
_FIELD_INSERT_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FIELD_INSERT_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_FIELD_INSERT_LINE}")"
_INDENTED_FIELD_INSERT_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_FIELD_INSERT_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_FIELD_INSERT_INDEX}" "${_INDENTED_FIELD_INSERT_SCRIPT}")"


#########################################################
# Update

_PATTERN="$("${ZARUBA_BIN}" list append '[]' "class[\t ]*Mem${_ZRB_PASCAL_APP_CRUD_ENTITY}Repo\(")"
_PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "def[ \t]+update\(")"
_PATTERN="$("${ZARUBA_BIN}" list append "${_PATTERN}" "mem_${_ZRB_SNAKE_APP_CRUD_ENTITY}\.updated_at[ \t]*=[ \t]*datetime\.datetime\.utcnow")"
_FIELD_UPDATE_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_FIELD_UPDATE_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_FIELD_UPDATE_LINE}")"
_INDENTED_FIELD_UPDATE_SCRIPT="$("${ZARUBA_BIN}" str fullIndent "${_FIELD_UPDATE_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" "${_FIELD_UPDATE_INDEX}" "${_INDENTED_FIELD_UPDATE_SCRIPT}")"


#########################################################
# Overwrite existing repo

chmod 755 "${_REPO_LOCATION}"
"${ZARUBA_BIN}" lines write "${_REPO_LOCATION}" "${_LINES}"

echo "Done updating mem repo"
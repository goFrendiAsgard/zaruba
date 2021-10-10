_PATTERN="[\t ]*(class[\t ]*${_ZRB_APP_CRUD_PASCAL_ENTITY}Data.*)"

_FIELD_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/schema_field_declaration.py")"
_FIELD_SCRIPT_LINES='["$1"]'
for _INDEX in $("${ZARUBA_HOME}/zaruba" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _FIELD_SCRIPT="    $("${ZARUBA_HOME}/zaruba" str replace "${_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _FIELD_SCRIPT_LINES="$("${ZARUBA_HOME}/zaruba" list append "${_FIELD_SCRIPT_LINES}" "${_FIELD_SCRIPT}")"
done

_ZRB_SCHEMA_FIELD_DECLARATION="$("${ZARUBA_HOME}/zaruba" list join "${_FIELD_SCRIPT_LINES}")"

_setReplacementMap "${_PATTERN}" "${_ZRB_SCHEMA_FIELD_DECLARATION}"
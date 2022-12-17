set -e
echo "Set test field"

_FIELD_SCRIPT_TEMPLATE="$(_readText "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/test_field.py")"
_FIELD_SCRIPT_LINES='["$1"]'
for _INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_BIN}" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_BIN}" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _FIELD_SCRIPT="    $("${ZARUBA_BIN}" str replace "${_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_FIELD_SCRIPT_LINES}" "${_FIELD_SCRIPT}")"
done

_ZRB_REPO_FIELD="$("${ZARUBA_BIN}" list join "${_FIELD_SCRIPT_LINES}")"

# test field pattern
_PATTERN="[\t ]*(dummy_ztpl_app_crud_entity_data[\t ]*=[\t ]*ZtplAppCrudEntityData\([\t ]*)"
_setReplacementMap "${_PATTERN}" "${_ZRB_REPO_FIELD}"

echo "Done setting test field"
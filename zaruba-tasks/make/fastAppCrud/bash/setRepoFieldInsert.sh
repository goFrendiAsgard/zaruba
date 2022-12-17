set -e
echo "Set repo field insert"

_FIELD_SCRIPT_TEMPLATE="$(_readText "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/repo_field_insert.py")"
_FIELD_SCRIPT_LINES='["$1"]'
for _INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_BIN}" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_BIN}" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _FIELD_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_FIELD_SCRIPT_LINES}" "${_FIELD_SCRIPT}")"
done

_ZRB_REPO_FIELD_INSERT="$("${ZARUBA_BIN}" list join "${_FIELD_SCRIPT_LINES}")"

# db and mem repo pattern
_PATTERN="[\t ]*(id[\t ]*=[\t ]*new_ztpl_app_crud_entity_id[\t ]*,[\t ]*)"
_setReplacementMap "${_PATTERN}" "${_ZRB_REPO_FIELD_INSERT}"

echo "Done setting repo field insert"
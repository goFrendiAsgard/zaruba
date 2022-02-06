set -e
echo "Set repo field update"

_PATTERN="[\t ]*(db_entity.updated_at[\t ]*=[\t ]datetime.datetime.utcnow\(.*)"

_FIELD_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/repo_field_update.py")"
_FIELD_SCRIPT_LINES='[]'
for _INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_BIN}" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_BIN}" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _FIELD_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_FIELD_SCRIPT_LINES}" "${_FIELD_SCRIPT}")"
done
_FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_FIELD_SCRIPT_LINES}" '$1')"


_ZRB_REPO_FIELD_UPDATE="$("${ZARUBA_BIN}" list join "${_FIELD_SCRIPT_LINES}")"

_setReplacementMap "${_PATTERN}" "${_ZRB_REPO_FIELD_UPDATE}"

echo "Done setting repo field update"
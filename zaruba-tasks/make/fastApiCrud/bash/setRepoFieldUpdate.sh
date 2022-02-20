set -e
echo "Set repo field update"

_DB_FIELD_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/repo_field_update_db.py")"
_DB_FIELD_SCRIPT_LINES='[]'
_MEM_FIELD_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/repo_field_update_mem.py")"
_MEM_FIELD_SCRIPT_LINES='[]'
for _INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_BIN}" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_BIN}" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _DB_FIELD_SCRIPT="$("${ZARUBA_BIN}" str replace "${_DB_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"
    _MEM_FIELD_SCRIPT="$("${ZARUBA_BIN}" str replace "${_MEM_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _DB_FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_DB_FIELD_SCRIPT_LINES}" "${_DB_FIELD_SCRIPT}")"
    _MEM_FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_MEM_FIELD_SCRIPT_LINES}" "${_MEM_FIELD_SCRIPT}")"
done
_DB_FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_DB_FIELD_SCRIPT_LINES}" '$1')"
_MEM_FIELD_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_MEM_FIELD_SCRIPT_LINES}" '$1')"

_ZRB_DB_REPO_FIELD_UPDATE="$("${ZARUBA_BIN}" list join "${_DB_FIELD_SCRIPT_LINES}")"
_ZRB_MEM_REPO_FIELD_UPDATE="$("${ZARUBA_BIN}" list join "${_MEM_FIELD_SCRIPT_LINES}")"

# db repo pattern
_DB_REPO_PATTERN="[\t ]*(db_ztpl_app_crud_entity.updated_at[\t ]*=[\t ]datetime.datetime.utcnow\(.*)"
_setReplacementMap "${_DB_REPO_PATTERN}" "${_ZRB_DB_REPO_FIELD_UPDATE}"

# mem repo pattern
_MEM_REPO_PATTERN="[\t ]*(mem_ztpl_app_crud_entity.updated_at[\t ]*=[\t ]datetime.datetime.utcnow\(.*)"
_setReplacementMap "${_MEM_REPO_PATTERN}" "${_ZRB_MEM_REPO_FIELD_UPDATE}"


echo "Done setting repo field update"
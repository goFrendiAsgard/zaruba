_PATTERN="[\t ]*(db_entity.updated_at[\t ]*=[\t ]datetime.datetime.utcnow\(.*)"

_FIELD_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiCrud/partials/repo_field_update.py")"
_FIELD_SCRIPT_LINES='[]'
for _INDEX in $("${ZARUBA_HOME}/zaruba" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _FIELD_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_FIELD_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _FIELD_SCRIPT_LINES="$("${ZARUBA_HOME}/zaruba" list append "${_FIELD_SCRIPT_LINES}" "${_FIELD_SCRIPT}")"
done
_FIELD_SCRIPT_LINES="$("${ZARUBA_HOME}/zaruba" list append "${_FIELD_SCRIPT_LINES}" '$1')"


_ZRB_REPO_FIELD_UPDATE="$("${ZARUBA_HOME}/zaruba" list join "${_FIELD_SCRIPT_LINES}")"

_setReplacementMap "${_PATTERN}" "${_ZRB_REPO_FIELD_UPDATE}"
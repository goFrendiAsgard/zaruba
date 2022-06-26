set -e
echo "Set jinja template"

_TABLE_HEADERS_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/jinja/table_header.html")"
_TABLE_HEADERS_SCRIPT_LINES='[]'
_TABLE_VALUES_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/jinja/table_values.html")"
_TABLE_VALUES_SCRIPT_LINES='[]'
_FORM_INPUTS_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/jinja/form_input.html")"
_FORM_INPUTS_SCRIPT_LINES='[]'

for _INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_ZRB_APP_CRUD_FIELDS}")
do
    _FIELD_NAME="$("${ZARUBA_BIN}" list get "${_ZRB_APP_CRUD_FIELDS}" "${_INDEX}")"
    _KEBAB_FIELD_NAME="$("${ZARUBA_BIN}" str toKebab "${_FIELD_NAME}")"
    _PASCAL_FIELD_NAME="$("${ZARUBA_BIN}" str toPascal "${_FIELD_NAME}")"
    _SNAKE_FIELD_NAME="$("${ZARUBA_BIN}" str toSnake "${_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_ZRB_REPLACEMENT_MAP}" "ztpl-app-crud-field" "${_KEBAB_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_LOCAL_REPLACEMENT_MAP}" "ZtplAppCrudField" "${_PASCAL_FIELD_NAME}")"
    _LOCAL_REPLACEMENT_MAP="$("${ZARUBA_BIN}" map set "${_LOCAL_REPLACEMENT_MAP}" "ztpl_app_crud_field" "${_SNAKE_FIELD_NAME}")"

    _TABLE_HEADERS_SCRIPT="$("${ZARUBA_BIN}" str replace "${_TABLE_HEADERS_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"
    _TABLE_VALUES_SCRIPT="$("${ZARUBA_BIN}" str replace "${_TABLE_VALUES_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"
    _FORM_INPUTS_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FORM_INPUTS_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _TABLE_HEADERS_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_TABLE_HEADERS_SCRIPT_LINES}" "${_TABLE_HEADERS_SCRIPT}")"
    _TABLE_VALUES_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_TABLE_VALUES_SCRIPT_LINES}" "${_TABLE_VALUES_SCRIPT}")"
    _FORM_INPUTS_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_FORM_INPUTS_SCRIPT_LINES}" "${_FORM_INPUTS_SCRIPT}")"
done

_ZRB_TABLE_HEADERS="$("${ZARUBA_BIN}" list join "${_TABLE_HEADERS_SCRIPT_LINES}")"
_ZRB_TABLE_VALUES="$("${ZARUBA_BIN}" list join "${_TABLE_VALUES_SCRIPT_LINES}")"
_ZRB_FORM_INPUTS="$("${ZARUBA_BIN}" list join "${_FORM_INPUTS_SCRIPT_LINES}")"

# table header pattern
_TABLE_HEADERS_PATTERN="[\t ]*<!-- CRUD column headers -->"
_setReplacementMap "${_TABLE_HEADERS_PATTERN}" "${_ZRB_TABLE_HEADERS}"

# table row pattern
_TABLE_VALUES_PATTERN="[\t ]*<!-- CRUD column values -->"
_setReplacementMap "${_TABLE_VALUES_PATTERN}" "${_ZRB_TABLE_VALUES}"

# form input pattern
_FORM_INPUTS_PATTERN="[\t ]*<!-- CRUD form inputs -->"
_setReplacementMap "${_FORM_INPUTS_PATTERN}" "${_ZRB_FORM_INPUTS}"

echo "Done setting jinja template"
set -e
echo "Set jinja template"

_COLUMN_HEADERS_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/jinja/column_header.html")"
_COLUMN_HEADERS_SCRIPT_LINES='[]'
_COLUMN_VALUES_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/jinja/column_values.html")"
_COLUMN_VALUES_SCRIPT_LINES='[]'
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

    _COLUMN_HEADERS_SCRIPT="$("${ZARUBA_BIN}" str replace "${_COLUMN_HEADERS_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"
    _COLUMN_VALUES_SCRIPT="$("${ZARUBA_BIN}" str replace "${_COLUMN_VALUES_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"
    _FORM_INPUTS_SCRIPT="$("${ZARUBA_BIN}" str replace "${_FORM_INPUTS_SCRIPT_TEMPLATE}" "${_LOCAL_REPLACEMENT_MAP}")"

    _COLUMN_HEADERS_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_COLUMN_HEADERS_SCRIPT_LINES}" "${_COLUMN_HEADERS_SCRIPT}")"
    _COLUMN_VALUES_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_COLUMN_VALUES_SCRIPT_LINES}" "${_COLUMN_VALUES_SCRIPT}")"
    _FORM_INPUTS_SCRIPT_LINES="$("${ZARUBA_BIN}" list append "${_FORM_INPUTS_SCRIPT_LINES}" "${_FORM_INPUTS_SCRIPT}")"
done

_ZRB_COLUMN_HEADERS="$("${ZARUBA_BIN}" list join "${_COLUMN_HEADERS_SCRIPT_LINES}")"
_ZRB_COLUMN_VALUES="$("${ZARUBA_BIN}" list join "${_COLUMN_VALUES_SCRIPT_LINES}")"
_ZRB_FORM_INPUTS="$("${ZARUBA_BIN}" list join "${_FORM_INPUTS_SCRIPT_LINES}")"

# column header pattern
_COLUMN_HEADERS_PATTERN="[\t ]*<!-- CRUD column headers -->"
_setReplacementMap "${_COLUMN_HEADERS_PATTERN}" "${_ZRB_COLUMN_HEADERS}"

# column value pattern
_COLUMN_VALUES_PATTERN="[\t ]*<!-- CRUD column values -->"
_setReplacementMap "${_COLUMN_VALUES_PATTERN}" "${_ZRB_COLUMN_VALUES}"

# form input pattern
_FORM_INPUTS_PATTERN="[\t ]*<!-- CRUD form inputs -->"
_setReplacementMap "${_FORM_INPUTS_PATTERN}" "${_ZRB_FORM_INPUTS}"

echo "Done setting jinja template"
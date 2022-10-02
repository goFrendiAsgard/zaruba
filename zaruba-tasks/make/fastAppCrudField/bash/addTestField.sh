set -e
echo "Updating test field"

_TEST_FIELD_SCRIPT_TEMPLATE="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrudField/partials/test_field.py")"
_TEST_FIELD_SCRIPT="$("${ZARUBA_BIN}" str replace "${_TEST_FIELD_SCRIPT_TEMPLATE}" "${_ZRB_REPLACEMENT_MAP}")"

#########################################################
# Read existing test
_TEST_LOCATION="${_ZRB_APP_DIRECTORY}/modules/${_ZRB_APP_MODULE_NAME}/${_ZRB_APP_CRUD_ENTITY}/test_${_ZRB_APP_CRUD_ENTITY}Service.py"
_LINES="$("${ZARUBA_BIN}" lines read "${_TEST_LOCATION}")"


#########################################################
# Mock entity data

_PATTERN="[\t ]*(mock_${_ZRB_SNAKE_APP_CRUD_ENTITY}_data[\t ]*=[\t ]*${_ZRB_PASCAL_APP_CRUD_ENTITY}Data\([\t ]*)"
_MOCK_ENTITY_DATA_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_MOCK_ENTITY_DATA_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_MOCK_ENTITY_DATA_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_MOCK_ENTITY_DATA_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_MOCK_ENTITY_DATA_LINE}")"
_INDENTED_MOCK_ENTITY_DATA_SCRIPT="    $("${ZARUBA_BIN}" str fullIndent "${_TEST_FIELD_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_MOCK_ENTITY_DATA_INDEX}" "${_INDENTED_MOCK_ENTITY_DATA_SCRIPT}")"


#########################################################
# Mock entity object

_PATTERN="[\t ]*(mock_${_ZRB_SNAKE_APP_CRUD_ENTITY}[\t ]*=[\t ]*${_ZRB_PASCAL_APP_CRUD_ENTITY}\([\t ]*)"
_MOCK_ENTITY_OBJECT_INDEX="$("${ZARUBA_BIN}" lines getIndex "${_LINES}" "${_PATTERN}")"
if [ "${_MOCK_ENTITY_OBJECT_INDEX}" = "-1" ]
then
    echo "Pattern not found: ${_PATTERN}"
    exit 1
fi

_MOCK_ENTITY_OBJECT_LINE="$("${ZARUBA_BIN}" list get "${_LINES}" "${_MOCK_ENTITY_OBJECT_INDEX}")"
_INDENTATION="$("${ZARUBA_BIN}" str getIndentation "${_MOCK_ENTITY_OBJECT_LINE}")"
_INDENTED_MOCK_ENTITY_OBJECT_SCRIPT="    $("${ZARUBA_BIN}" str fullIndent "${_TEST_FIELD_SCRIPT}" "${_INDENTATION}")"
_LINES="$("${ZARUBA_BIN}" lines insertAfter "${_LINES}" "${_MOCK_ENTITY_OBJECT_INDEX}" "${_INDENTED_MOCK_ENTITY_OBJECT_SCRIPT}")"


#########################################################
# Overwrite existing test

chmod 755 "${_SCHEMA_LOCATION}"
"${ZARUBA_BIN}" lines write "${_TEST_LOCATION}" "${_LINES}"

echo "Done updating test field"
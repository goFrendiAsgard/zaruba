set -e
echo "Import internal route handler"

_IMPORT_ROUTE_HANDLER_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_route_handler.py")"
_IMPORT_ROUTE_HANDLER_SCRIPT="$("${ZARUBA_BIN}" str replace "${_IMPORT_ROUTE_HANDLER_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_CONTROLLER_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/route.py"

_LINES="$("${ZARUBA_BIN}" lines read "${_CONTROLLER_FILE_LOCATION}")"

# insert import
_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_LINES}" 0 "${_IMPORT_ROUTE_HANDLER_SCRIPT}")"

chmod 755 "${_CONTROLLER_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_CONTROLLER_FILE_LOCATION}" "${_LINES}"

echo "Done importing internal route handler"
set -e
echo "Importing CRUD route handler"

# module/<module_name>/route.py
_importCrudRouteHandler() {
    _FILE_PATH="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/route.py"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_route_handler.py")"
    _insertPartialBefore "${_FILE_PATH}" "${_NEW_CONTENT}" 0
    chmod 755 "${_FILE_PATH}"
}

_importCrudRouteHandler

echo "Done importing CRUD route handler"
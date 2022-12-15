set -e
echo "Registering migration"

# module/<module_name>/_alembic/env.py
_registerMigration() {
    _FILE_PATH="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/_alembic/env.py"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_entity.py")"
    _insertPartialBefore "${_FILE_PATH}" "${_NEW_CONTENT}" 1
    chmod 755 "${_FILE_PATH}"
}

_registerMigration

echo "Done registering migration"
set -e
echo "Registering migration"

_IMPORT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_entity.py")"
_IMPORT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_IMPORT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"


####################################################################
## module/<module-name>/_alembic/env.py
####################################################################

####################################################################
# Read existing alembic env

_ENV_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/module/${_ZRB_SNAKE_APP_MODULE_NAME}/_alembic/env.py"
_ENV_LINES="$("${ZARUBA_BIN}" lines read "${_ENV_FILE_LOCATION}")"

_ENV_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_ENV_LINES}" "${_IMPORT_SCRIPT}" --index=1)"

####################################################################
# Overwrite existing alembic env

chmod 755 "${_ENV_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_ENV_LINES}" "${_ENV_FILE_LOCATION}"

echo "Done registering migration"
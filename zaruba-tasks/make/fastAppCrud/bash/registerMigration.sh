set -e
echo "Registering migration"

_IMPORT_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppCrud/partials/import_entity.py")"
_IMPORT_SCRIPT="$("${ZARUBA_BIN}" str replace "${_IMPORT_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"


####################################################################
## alembic/env.py
####################################################################

####################################################################
# Read existing alembic env

_ENV_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/alembic/env.py"
_ENV_LINES="$("${ZARUBA_BIN}" lines read "${_ENV_FILE_LOCATION}")"


_ENV_LINES="$("${ZARUBA_BIN}" lines insertBefore "${_ENV_LINES}" 1 "${_IMPORT_SCRIPT}")"

####################################################################
# Overwrite existing alembic env

chmod 755 "${_ENV_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_ENV_FILE_LOCATION}" "${_ENV_LINES}"

echo "Done registering migration"
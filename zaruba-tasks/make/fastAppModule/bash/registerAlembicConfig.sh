echo "Registering alembic config"

_ALEMBIC_CONFIG_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/alembic.ini")"
_ALEMBIC_CONFIG_SCRIPT="$("${ZARUBA_BIN}" str replace "${_ALEMBIC_CONFIG_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

#########################################################
# Read existing env

_ALEMBIC_CONFIG_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/alembic.ini"
_ALEMBIC_CONFIG_LINES="$("${ZARUBA_BIN}" lines read "${_ALEMBIC_CONFIG_FILE_LOCATION}")"

#########################################################
# Add new env

_ALEMBIC_CONFIG_LINES="$("${ZARUBA_BIN}" list append "${_ALEMBIC_CONFIG_LINES}" "${_ALEMBIC_CONFIG_SCRIPT}")"

#########################################################
# Overwrite existing repo

chmod 755 "${_ALEMBIC_CONFIG_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_ALEMBIC_CONFIG_FILE_LOCATION}" "${_ALEMBIC_CONFIG_LINES}"

echo "Done registering alembic config"
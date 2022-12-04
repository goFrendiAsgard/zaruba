echo "Registering migrate command"

_MIGRATE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/migrate.sh")"
_MIGRATE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_MIGRATE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

#########################################################
# Read existing env

_MIGRATE_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/migrate.sh"
_MIGRATE_LINES="$("${ZARUBA_BIN}" lines read "${_MIGRATE_FILE_LOCATION}")"

#########################################################
# Add new env

_MIGRATE_LINES="$("${ZARUBA_BIN}" list append "${_MIGRATE_LINES}" "${_MIGRATE_SCRIPT}")"

#########################################################
# Overwrite existing repo

chmod 755 "${_MIGRATE_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_MIGRATE_FILE_LOCATION}" "${_MIGRATE_LINES}"

echo "Done registering migrate command"
echo "Registering create migration command"

_CREATE_MIGRATION_SCRIPT="$(_readText "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/create-migration.sh")"
_CREATE_MIGRATION_SCRIPT="$("${ZARUBA_BIN}" str replace "${_CREATE_MIGRATION_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

#########################################################
# Read existing env

_CREATE_MIGRATION_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/create-migration.sh"
_CREATE_MIGRATION_LINES="$(_readLines "${_CREATE_MIGRATION_FILE_LOCATION}")"

#########################################################
# Add new env

_CREATE_MIGRATION_LINES="$("${ZARUBA_BIN}" list append "${_CREATE_MIGRATION_LINES}" "${_CREATE_MIGRATION_SCRIPT}")"

#########################################################
# Overwrite existing repo

chmod 755 "${_CREATE_MIGRATION_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_CREATE_MIGRATION_LINES}" "${_CREATE_MIGRATION_FILE_LOCATION}"

echo "Done registering create migration command"
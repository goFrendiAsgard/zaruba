echo "Registering create migration command"


_registerCreateMigrationCommand() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/create-migration.sh"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/create-migration.sh")"
    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" -1
    chmod 755 "${_DESTINATION}"
}

_registerCreateMigrationCommand

echo "Done registering create migration command"
echo "Registering migrate command"

_registerMigrateCommand() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/migrate.sh"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/migrate.sh")"
    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" -1
    chmod 755 "${_DESTINATION}"
}

_registerMigrateCommand

echo "Done registering migrate command"
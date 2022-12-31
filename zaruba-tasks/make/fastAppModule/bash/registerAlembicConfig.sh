echo "Registering alembic config"

_registerAlembicConfig() {
    _DESTINATION="${ZARUBA_PROJECT_DIR}/${_ZRB_APP_DIRECTORY}/alembic.ini"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/alembic.ini")"
    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" -1
    chmod 755 "${_DESTINATION}"
}

_registerAlembicConfig

echo "Done registering alembic config"
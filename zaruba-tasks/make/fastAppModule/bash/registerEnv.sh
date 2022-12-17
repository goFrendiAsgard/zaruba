echo "Registering env"

_registerEnv() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/template.env"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/template.env")"
    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" -1
    chmod 755 "${_DESTINATION}"
}

_registerEnv
"${ZARUBA_BIN}" task syncEnv "start${_ZRB_PASCAL_APP_NAME}Container"

echo "Done registering env"
echo "Registering module"

_importModule() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/main.py"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/import_module.py")"
    _insertPartialBefore "${_DESTINATION}" "${_NEW_CONTENT}" 0
    chmod 755 "${_DESTINATION}"
}

_registerModule() {
    _DESTINATION="${_ZRB_APP_DIRECTORY}/main.py"
    _NEW_CONTENT="$(_getPartialContent "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/load_module.py")"
    _insertPartialAfter "${_DESTINATION}" "${_NEW_CONTENT}" -1
    chmod 755 "${_DESTINATION}"
}

echo "Importing module"
_importModule
echo "Registering module"
_registerModule

echo "Done registering module"
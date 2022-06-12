echo "Registering module"

_IMPORT_MODULE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiModule/partials/import_module.py")"
_IMPORT_MODULE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_IMPORT_MODULE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_LOAD_MODULE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiModule/partials/load_module.py")"
_LOAD_MODULE_SCRIPT="$("${ZARUBA_BIN}" str replace "${_LOAD_MODULE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_MAIN_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/main.py"
_MAIN_SCRIPT="$(cat "${_MAIN_FILE_LOCATION}")"

_MAIN_LINES="$("${ZARUBA_BIN}" list append '[]' "${_IMPORT_MODULE_SCRIPT}")"
_MAIN_LINES="$("${ZARUBA_BIN}" list append "${_MAIN_LINES}" "${_MAIN_SCRIPT}")"
_MAIN_LINES="$("${ZARUBA_BIN}" list append "${_MAIN_LINES}" "${_LOAD_MODULE_SCRIPT}")"

chmod 755 "${_MAIN_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_MAIN_FILE_LOCATION}" "${_MAIN_LINES}"

echo "Done registering module"
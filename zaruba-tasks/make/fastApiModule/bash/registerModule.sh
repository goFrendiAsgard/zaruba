_IMPORT_MODULE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiModule/partials/import_module.py")"
_IMPORT_MODULE_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_MODULE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_LOAD_MODULE_SCRIPT="$(cat "${ZARUBA_HOME}/zaruba-tasks/make/fastApiModule/partials/load_module.py")"
_LOAD_MODULE_SCRIPT="$("${ZARUBA_HOME}/zaruba" str replace "${_LOAD_MODULE_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

_MAIN_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/main.py"
_MAIN_SCRIPT="$(cat "${_MAIN_FILE_LOCATION}")"

_LINES="$("${ZARUBA_HOME}/zaruba" list append '[]' "${_IMPORT_MODULE_SCRIPT}")"
_LINES="$("${ZARUBA_HOME}/zaruba" list append "${_LINES}" "${_MAIN_SCRIPT}")"
_LINES="$("${ZARUBA_HOME}/zaruba" list append "${_LINES}" "${_LOAD_MODULE_SCRIPT}")"

chmod 755 "${_MAIN_FILE_LOCATION}"
"${ZARUBA_HOME}/zaruba" lines write "${_MAIN_FILE_LOCATION}" "${_LINES}"
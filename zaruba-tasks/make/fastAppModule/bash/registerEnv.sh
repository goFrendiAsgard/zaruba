echo "Registering env"

_ENV_SCRIPT="$(_readText "${ZARUBA_HOME}/zaruba-tasks/make/fastAppModule/partials/template.env")"
_ENV_SCRIPT="$("${ZARUBA_BIN}" str replace "${_ENV_SCRIPT}" "${_ZRB_REPLACEMENT_MAP}" )"

#########################################################
# Read existing env

_ENV_FILE_LOCATION="${_ZRB_APP_DIRECTORY}/template.env"
_ENV_LINES="$(_readLines "${_ENV_FILE_LOCATION}")"

#########################################################
# Add new env

_ENV_LINES="$("${ZARUBA_BIN}" list append "${_ENV_LINES}" "${_ENV_SCRIPT}")"

#########################################################
# Overwrite existing repo

chmod 755 "${_ENV_FILE_LOCATION}"
"${ZARUBA_BIN}" lines write "${_ENV_LINES}" "${_ENV_FILE_LOCATION}"

#########################################################
# sync env

"${ZARUBA_BIN}" task syncEnv "start${_ZRB_PASCAL_APP_NAME}Container"

echo "Done registering env"
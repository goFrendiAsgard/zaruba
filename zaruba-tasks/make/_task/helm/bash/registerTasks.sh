_PROJECT_FILE_NAME="${1}"
_APP_NAME="${2}"

_PASCAL_APP_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_APP_NAME}")" 

# helmInstall
_registerTask "${_PROJECT_FILE_NAME}" "helmInstall" "helmInstall${_PASCAL_APP_NAME}"
# helmUninstall
_registerTask "${_PROJECT_FILE_NAME}" "helmUninstall" "helmUninstall${_PASCAL_APP_NAME}"
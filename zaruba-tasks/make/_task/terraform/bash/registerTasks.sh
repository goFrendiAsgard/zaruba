_PROJECT_FILE_NAME="${1}"
_APP_NAME="${2}"

_PASCAL_APP_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_APP_NAME}")" 

# terraformApply
_registerTask "${_PROJECT_FILE_NAME}" "terraformApply" "terraformApply${_PASCAL_APP_NAME}"
# terraformDestroy
_registerTask "${_PROJECT_FILE_NAME}" "terraformDestroy" "terraformDestroy${_PASCAL_APP_NAME}"
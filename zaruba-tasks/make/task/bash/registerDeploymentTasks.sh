_PROJECT_FILE_NAME="${1}"
_DEPLOYMENT_NAME="${2}"

_PASCAL_DEPLOYMENT_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_DEPLOYMENT_NAME}")" 

# deploy
_registerTask "${_PROJECT_FILE_NAME}" "deploy" "deploy${_PASCAL_DEPLOYMENT_NAME}"
# destroy
_registerTask "${_PROJECT_FILE_NAME}" "destroy" "destroy${_PASCAL_DEPLOYMENT_NAME}"
echo "Registering deployment tasks"

_PROJECT_FILE_NAME="${1}"
_DEPLOYMENT_NAME="${2}"

_PASCAL_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" str toPascal "${_DEPLOYMENT_NAME}")" 

_registerTask "${_PROJECT_FILE_NAME}" "prepareDeployments" "prepare${_PASCAL_DEPLOYMENT_NAME}Deployment"
_registerTask "${_PROJECT_FILE_NAME}" "deploy" "deploy${_PASCAL_DEPLOYMENT_NAME}"
_registerTask "${_PROJECT_FILE_NAME}" "destroy" "destroy${_PASCAL_DEPLOYMENT_NAME}"

echo "Done registering deploymentr tasks"
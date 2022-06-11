echo "Registering deployment tasks"

_PROJECT_FILE_NAME="${1}"
_DEPLOYMENT_NAME="${2}"

_PASCAL_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" str toPascal "${_DEPLOYMENT_NAME}")" 

_addTaskDependency "prepareDeployments" "prepare${_PASCAL_DEPLOYMENT_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "deploy" "deploy${_PASCAL_DEPLOYMENT_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "destroy" "destroy${_PASCAL_DEPLOYMENT_NAME}" 1 ${_PROJECT_FILE_NAME}

echo "Done registering deployment tasks"
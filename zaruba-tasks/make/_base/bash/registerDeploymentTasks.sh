echo "Registering deployment tasks"

_DEPLOYMENT_NAME="${1}"

_PASCAL_DEPLOYMENT_NAME="$("${ZARUBA_BIN}" str toPascal "${_DEPLOYMENT_NAME}")" 

_addTaskDependency "prepareDeployments" "prepare${_PASCAL_DEPLOYMENT_NAME}" 1
_addTaskDependency "deploy" "deploy${_PASCAL_DEPLOYMENT_NAME}" 1
_addTaskDependency "destroy" "destroy${_PASCAL_DEPLOYMENT_NAME}" 1

echo "Done registering deployment tasks"
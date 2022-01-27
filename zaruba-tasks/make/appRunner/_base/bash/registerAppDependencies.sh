
_PROJECT_FILE_NAME="${1}"
_APP_NAME="${2}"
_APP_DEPENDENCIES="${3}"

_PASCAL_APP_NAME="$("${ZARUBA_BIN}" str toPascal "${_APP_NAME}")" 

echo "Registering start${_PASCAL_APP_NAME}Container dependencies"
_CONTAINER_TASK_FILE_NAME="zaruba-tasks/${_APP_NAME}/tasks.container.yaml"
for _APP_DEPENDENCY_INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_APP_DEPENDENCIES}")
do
    _APP_DEPENDENCY_NAME="$("${ZARUBA_BIN}" list get "${_APP_DEPENDENCIES}" "${_APP_DEPENDENCY_INDEX}")"
    _PASCAL_APP_DEPENDENCY_NAME="$("${ZARUBA_BIN}" str toPascal "${_APP_DEPENDENCY_NAME}")" 
    if [ "$("${ZARUBA_BIN}" task isExist "${_PROJECT_FILE_NAME}" "start${_PASCAL_APP_NAME}Container")" = 1 ]
    then
        _addTaskDependency "${_PROJECT_FILE_NAME}" "start${_PASCAL_APP_NAME}Container" "start${_PASCAL_APP_DEPENDENCY_NAME}Container" 0
    else
        _addTaskDependency "${_PROJECT_FILE_NAME}" "start${_PASCAL_APP_NAME}Container" "start${_PASCAL_APP_DEPENDENCY_NAME}" 0
    fi
done


echo "Registering start${_PASCAL_APP_NAME} dependencies"
_TASK_FILE_NAME="zaruba-tasks/${_APP_NAME}/tasks.yaml"
for _APP_DEPENDENCY_INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_APP_DEPENDENCIES}")
do
    _APP_DEPENDENCY_NAME="$("${ZARUBA_BIN}" list get "${_APP_DEPENDENCIES}" "${_APP_DEPENDENCY_INDEX}")"
    _PASCAL_APP_DEPENDENCY_NAME="$("${ZARUBA_BIN}" str toPascal "${_APP_DEPENDENCY_NAME}")" 
    _addTaskDependency "${_PROJECT_FILE_NAME}" "start${_PASCAL_APP_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}" 0
done
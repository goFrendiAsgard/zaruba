
_PROJECT_FILE_NAME="${1}"
_APP_NAME="${2}"
_APP_DEPENDENCIES="${3}"
_CONTAINER_PREPARE_APP_RUNNER_TASK_NAME="${4}"
_NATIVE_PREPARE_APP_RUNNER_TASK_NAME="${5}"

for _APP_DEPENDENCY_INDEX in $("${ZARUBA_BIN}" list rangeIndex "${_APP_DEPENDENCIES}")
do
    _APP_DEPENDENCY_NAME="$("${ZARUBA_BIN}" list get "${_APP_DEPENDENCIES}" "${_APP_DEPENDENCY_INDEX}")"
    _PASCAL_APP_DEPENDENCY_NAME="$("${ZARUBA_BIN}" str toPascal "${_APP_DEPENDENCY_NAME}")" 
    # container app runner
    if [ ! -z "${_CONTAINER_PREPARE_APP_RUNNER_TASK_NAME}" ]
    then
        if [ "$("${ZARUBA_BIN}" task isExist "${_PROJECT_FILE_NAME}" "start${_PASCAL_APP_DEPENCENCY_NAME}Container")" = 1 ]
        then
            _addTaskDependency "${_PROJECT_FILE_NAME}" "${_CONTAINER_PREPARE_APP_RUNNER_TASK_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}Container" 0
        else
            _addTaskDependency "${_PROJECT_FILE_NAME}" "${_CONTAINER_PREPARE_APP_RUNNER_TASK_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}" 0
        fi
    fi
    # native app runner
    if [ ! -z "${_NATIVE_PREPARE_APP_RUNNER_TASK_NAME}" ]
    then
        _addTaskDependency "${_PROJECT_FILE_NAME}" "${_NATIVE_PREPARE_APP_RUNNER_TASK_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}" 0
    fi
done
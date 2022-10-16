
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
        if [ "$("${ZARUBA_BIN}" task isExist "start${_PASCAL_APP_DEPENCENCY_NAME}Container" "${_PROJECT_FILE_NAME}")" = 1 ]
        then
            _addTaskDependency "${_CONTAINER_PREPARE_APP_RUNNER_TASK_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}Container" 0 ${_PROJECT_FILE_NAME}
        else
            _addTaskDependency "${_CONTAINER_PREPARE_APP_RUNNER_TASK_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}" 0 ${_PROJECT_FILE_NAME}
        fi
    fi
    # native app runner
    if [ ! -z "${_NATIVE_PREPARE_APP_RUNNER_TASK_NAME}" ]
    then
        _addTaskDependency "${_NATIVE_PREPARE_APP_RUNNER_TASK_NAME}" "start${_PASCAL_APP_DEPENDENCY_NAME}" 0 ${_PROJECT_FILE_NAME}
    fi
done
echo "Registering app runner tasks"

_SCRIPT_FILE_NAME="${1}"
_APP_NAME="${2}"

_PASCAL_APP_NAME="$("${ZARUBA_BIN}" str toPascal "${_APP_NAME}")" 

_addTaskDependency "${_SCRIPT_FILE_NAME}" "prepare" "prepare${_PASCAL_APP_NAME}" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "test" "test${_PASCAL_APP_NAME}" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "migrate" "migrate${_PASCAL_APP_NAME}" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "start" "start${_PASCAL_APP_NAME}" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "startContainers" "start${_PASCAL_APP_NAME}Container" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "run" "run${_PASCAL_APP_NAME}" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "runContainers" "run${_PASCAL_APP_NAME}Container" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "stopContainers" "stop${_PASCAL_APP_NAME}Container" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "removeContainers" "remove${_PASCAL_APP_NAME}Container" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "buildImages" "build${_PASCAL_APP_NAME}Image" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "pushImages" "push${_PASCAL_APP_NAME}Image" 1 "${SCRIPT_FILE_NAME}"
_addTaskDependency "${_SCRIPT_FILE_NAME}" "pullImages" "pull${_PASCAL_APP_NAME}Image" 1 "${SCRIPT_FILE_NAME}"

echo "Done registering app runner tasks"
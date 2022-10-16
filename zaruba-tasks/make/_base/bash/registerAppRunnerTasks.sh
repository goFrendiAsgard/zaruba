echo "Registering app runner tasks"

_PROJECT_FILE_NAME="${1}"
_APP_NAME="${2}"

_PASCAL_APP_NAME="$("${ZARUBA_BIN}" str toPascal "${_APP_NAME}")" 

_addTaskDependency "prepare" "prepare${_PASCAL_APP_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "test" "test${_PASCAL_APP_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "migrate" "migrate${_PASCAL_APP_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "start" "start${_PASCAL_APP_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "startContainers" "start${_PASCAL_APP_NAME}Container" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "run" "run${_PASCAL_APP_NAME}" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "runContainers" "run${_PASCAL_APP_NAME}Container" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "stopContainers" "stop${_PASCAL_APP_NAME}Container" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "removeContainers" "remove${_PASCAL_APP_NAME}Container" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "buildImages" "build${_PASCAL_APP_NAME}Image" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "pushImages" "push${_PASCAL_APP_NAME}Image" 1 ${_PROJECT_FILE_NAME}
_addTaskDependency "pullImages" "pull${_PASCAL_APP_NAME}Image" 1 ${_PROJECT_FILE_NAME}

echo "Done registering app runner tasks"
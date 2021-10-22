_PROJECT_FILE_NAME="${1}"
_INDEX_FILE_NAME="${2}"
_APP_NAME="${3}"

_PASCAL_APP_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_APP_NAME}")" 
# prepare
_registerTask "${_PROJECT_FILE_NAME}" "prepare" "prepare${_PASCAL_APP_NAME}"
# test
_registerTask "${_PROJECT_FILE_NAME}" "test" "test${_PASCAL_APP_NAME}"
# start
_registerTask "${_PROJECT_FILE_NAME}" "start" "start${_PASCAL_APP_NAME}"
# startContainers
_registerTask "${_PROJECT_FILE_NAME}" "startContainers" "start${_PASCAL_APP_NAME}Container"
# run
_registerTask "${_PROJECT_FILE_NAME}" "run" "run${_PASCAL_APP_NAME}"
# runContainers
_registerTask "${_PROJECT_FILE_NAME}" "runContainers" "run${_PASCAL_APP_NAME}Container"
# stopContainers
_registerTask "${_PROJECT_FILE_NAME}" "stopContainers" "stop${_PASCAL_APP_NAME}Container"
# removeContainers
_registerTask "${_PROJECT_FILE_NAME}" "removeContainers" "remove${_PASCAL_APP_NAME}Container"
# buildImages
_registerTask "${_PROJECT_FILE_NAME}" "buildImages" "build${_PASCAL_APP_NAME}Image"
# pushImages
_registerTask "${_PROJECT_FILE_NAME}" "pushImages" "push${_PASCAL_APP_NAME}Image"
# pullImages
_registerTask "${_PROJECT_FILE_NAME}" "pullImages" "pull${_PASCAL_APP_NAME}Image"
# helmInstall
_registerTask "${_PROJECT_FILE_NAME}" "helmInstall" "helmInstall${_PASCAL_APP_NAME}"
# helmUninstall
_registerTask "${_PROJECT_FILE_NAME}" "helmUninstall" "helmUninstall${_PASCAL_APP_NAME}"
# terraformApply
_registerTask "${_PROJECT_FILE_NAME}" "terraformApply" "terraformApply${_PASCAL_APP_NAME}"
# terraformDestroy
_registerTask "${_PROJECT_FILE_NAME}" "terraformDestroy" "terraformDestroy${_PASCAL_APP_NAME}"
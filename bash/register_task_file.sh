# USAGE: register_task_file <task-file-name> <service-name>
register_task_file() {
    _TASK_FILE_NAME="${1}"
    _SERVICE_NAME="${2}"
    _PASCAL_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToPascal "${_SERVICE_NAME}")" 

    "${ZARUBA_HOME}/zaruba" project include "./main.zaruba.yaml" "${_TASK_FILE_NAME}"
    "${ZARUBA_HOME}/zaruba" syncProjectEnvFiles "./main.zaruba.yaml"

    # buildImage
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "build${_PASCAL_SERVICE_NAME}Image")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "buildImage"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "buildImage" "[\"build${_PASCAL_SERVICE_NAME}Image\"]"
    fi

    # pullImage
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "pull${_PASCAL_SERVICE_NAME}Image")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "pullImage"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "pullImage" "[\"pull${_PASCAL_SERVICE_NAME}Image\"]"
    fi

    # pushImage
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "push${_PASCAL_SERVICE_NAME}Image")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "pushImage"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "pushImage" "[\"push${_PASCAL_SERVICE_NAME}Image\"]"
    fi

    # run
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "run${_PASCAL_SERVICE_NAME}")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "run"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "run" "[\"run${_PASCAL_SERVICE_NAME}\"]"
        # runContainer (in case of service's runContainer task doesn't exist)
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "runContainer"
        if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "run${_PASCAL_SERVICE_NAME}Container")" -eq 0 ]
        then
            "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "runContainer"
            "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "runContainer" "[\"run${_PASCAL_SERVICE_NAME}\"]"
        fi
    fi

    # runContainer
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "run${_PASCAL_SERVICE_NAME}Container")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "runContainer"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "runContainer" "[\"run${_PASCAL_SERVICE_NAME}Container\"]"
    fi 

    # stopContainer
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "stop${_PASCAL_SERVICE_NAME}Container")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "stopContainer"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "stopContainer" "[\"stop${_PASCAL_SERVICE_NAME}Container\"]"
    fi

    # removeContainer
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "remove${_PASCAL_SERVICE_NAME}Container")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "removeContainer"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "removeContainer" "[\"remove${_PASCAL_SERVICE_NAME}Container\"]"
    fi

    # helmInstall
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "helmInstall${_PASCAL_SERVICE_NAME}")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "helmInstall"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "helmInstall" "[\"helmInstall${_PASCAL_SERVICE_NAME}\"]"
    fi

    # helmUninstall
    if [ "$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "helmUninstall${_PASCAL_SERVICE_NAME}")" -eq 1 ]
    then
        "${ZARUBA_HOME}/zaruba" ensureTaskExist "./main.zaruba.yaml" "helmUninstall"
        "${ZARUBA_HOME}/zaruba" task addDependency "./main.zaruba.yaml" "helmUninstall" "[\"helmUninstall${_PASCAL_SERVICE_NAME}\"]"
    fi
}
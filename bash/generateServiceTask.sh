. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generatorUtil.sh
. ${ZARUBA_HOME}/bash/registerTaskFile.sh

# USAGE generateServiceTask <template-location> <service-location> <service-name> <image-name> <container-name> <service-start-command> <service-runner-version> <service-ports> <service-envs> <dependencies> <replacement-map> <register-runner>
generateServiceTask() {
    _TEMPLATE_LOCATION="${1}"
    _SERVICE_LOCATION="${2}"
    _SERVICE_NAME="${3}"
    _IMAGE_NAME="${4}"
    _CONTAINER_NAME="${5}"
    _SERVICE_START_COMMAND="${6}"
    _SERVICE_RUNNER_VERSION="${7}"
    _SERVICE_PORTS="${8}"
    _SERVICE_ENVS="${9}"
    _DEPENDENCIES="${10}"
    _REPLACEMENT_MAP="${11}"
    _REGISTER_RUNNER="${12}"

    _SERVICE_NAME="$(getServiceName "${_SERVICE_NAME}" "${_SERVICE_LOCATION}")"
    _IMAGE_NAME="$(getServiceImageName "${_IMAGE_NAME}" "${_SERVICE_NAME}")"
    _CONTAINER_NAME="$(getServiceContainerName "${_CONTAINER_NAME}" "${_SERVICE_NAME}")"

    _PASCAL_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_SERVICE_NAME}")"
    _KEBAB_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" str toKebab "${_SERVICE_NAME}")"
    _SNAKE_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" str toSnake "${_SERVICE_NAME}")"
    _UPPER_SNAKE_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" str toUpper "${_SNAKE_SERVICE_NAME}")"

    _TASK_EXIST="$("${ZARUBA_HOME}/zaruba" task isExist "./main.zaruba.yaml" "run${_PASCAL_SERVICE_NAME}")"
    if [ "${_TASK_EXIST}" -eq 1 ]
    then
        echo "service task already exist: run${_PASCAL_SERVICE_NAME}"
        return
    fi

    if [ "$("${ZARUBA_HOME}/zaruba" map validate "${_SERVICE_ENVS}")" -eq 0 ]
    then
        echo "env ${_SERVICE_ENVS} is not a valid map, apply default value"
        _SERVICE_ENVS='{}'
    fi 

    if [ "$("${ZARUBA_HOME}/zaruba" list validate "${_SERVICE_PORTS}")" -eq 0 ]
    then
        echo "ports ${_SERVICE_PORTS} is not a valid list, apply default value"
        _SERVICE_PORTS='[]'
    fi

    if [ "$("${ZARUBA_HOME}/zaruba" list validate "${_DEPENDENCIES}")" -eq 0 ]
    then
        echo "dependencies ${_DEPENDENCIES} is not a valid list, apply default value"
        _DEPENDENCIES='[]'
    fi

    _DESTINATION="."
    _TASK_FILE_NAME="${_DESTINATION}/zaruba-tasks/${_SERVICE_NAME}/task.zaruba.yaml"
    if [ -f "${_TASK_FILE_NAME}" ]
    then
        echo "file already exist: ${_TASK_FILE_NAME}"
        exit 1
    fi

    _DEFAULT_PORT_LIST="$("${ZARUBA_HOME}/zaruba" path getPortConfig "${_SERVICE_LOCATION}")"
    _DEFAULT_PORT_CONFIG="$("${ZARUBA_HOME}/zaruba" list join "${_DEFAULT_PORT_LIST}" "\n")"

    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" map set "${_REPLACEMENT_MAP}" \
        "zarubaImageName" "${_IMAGE_NAME}" \
        "zarubaContainerName" "${_CONTAINER_NAME}" \
        "zarubaServiceName" "${_SERVICE_NAME}" \
        "ZarubaServiceName" "${_PASCAL_SERVICE_NAME}" \
        "zaruba-service-name" "${_KEBAB_SERVICE_NAME}" \
        "ZARUBA_SERVICE_NAME" "${_UPPER_SNAKE_SERVICE_NAME}" \
        "zarubaStartCommand" "${_SERVICE_START_COMMAND}" \
        "zarubaServiceLocation" "$("${ZARUBA_HOME}/zaruba" path getRelativePath "$(dirname "${_TASK_FILE_NAME}")" "${_SERVICE_LOCATION}")" \
        "zarubaRunnerVersion" "${_SERVICE_RUNNER_VERSION}" \
        "zarubaDefaultPortConfig" "${_DEFAULT_PORT_CONFIG}" \
    ) 

    "${ZARUBA_HOME}/zaruba" util generate "${_TEMPLATE_LOCATION}" "${_DESTINATION}" "${_REPLACEMENT_MAP}"

    registerTaskFile "${_TASK_FILE_NAME}" "${_SERVICE_NAME}" "${_REGISTER_RUNNER}"

    "${ZARUBA_HOME}/zaruba" task addDependency ./main.zaruba.yaml "run${_PASCAL_SERVICE_NAME}" "${_DEPENDENCIES}"
    "${ZARUBA_HOME}/zaruba" task setEnv ./main.zaruba.yaml "run${_PASCAL_SERVICE_NAME}" "${_SERVICE_ENVS}"

            
    if [ "$("${ZARUBA_HOME}/zaruba" list length "${_SERVICE_PORTS}")" -gt 0 ]
    then
        _PORT_CONFIG_VALUE="$("${ZARUBA_HOME}/zaruba" list join "${_SERVICE_PORTS}" )"
        _PORT_CONFIG="$("${ZARUBA_HOME}/zaruba" map set "{}" "ports" "$_PORT_CONFIG_VALUE" )"
        "${ZARUBA_HOME}/zaruba" task setConfig ./main.zaruba.yaml "run${_PASCAL_SERVICE_NAME}" "${_PORT_CONFIG}"
    fi

}
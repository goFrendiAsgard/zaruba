. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/register_task_file.sh

# USAGE generate_service_task <template-location> <service-location> <service-name> <image-name> <container-name> <service-start-command> <service-runner-version> <service-ports> <service-envs> <dependencies> <replacement-map>
generate_service_task() {
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

    _DEFAULT_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" getServiceName "${_SERVICE_LOCATION}")"
    _SERVICE_NAME="$(get_value_or_default "${_SERVICE_NAME}" "${_DEFAULT_SERVICE_NAME}")"

    _DEFAULT_IMAGE_NAME="$("${ZARUBA_HOME}/zaruba" strToKebab "${_SERVICE_NAME}")"
    _IMAGE_NAME="$(get_value_or_default "${_IMAGE_NAME}" "${_DEFAULT_IMAGE_NAME}")"

    _DEFAULT_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" strToCamel "${_SERVICE_NAME}")"
    _CONTAINER_NAME="$(get_value_or_default "${_CONTAINER_NAME}" "${_DEFAULT_CONTAINER_NAME}")"

    _PASCAL_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToPascal "${_SERVICE_NAME}")"
    _KEBAB_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToKebab "${_SERVICE_NAME}")"
    _SNAKE_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToSnake "${_SERVICE_NAME}")"
    _UPPER_SNAKE_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToUpper "${_SNAKE_SERVICE_NAME}")"

    _TASK_EXIST="$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "run${_PASCAL_SERVICE_NAME}")"
    if [ "${_TASK_EXIST}" -eq 1 ]
    then
        echo "service task already exist: run${_PASCAL_SERVICE_NAME}"
        return
    fi

    if [ "$("${ZARUBA_HOME}/zaruba" isValidMap "${_SERVICE_ENVS}")" -eq 0 ]
    then
        echo "env ${_SERVICE_ENVS} is not a valid map, apply default value"
        _SERVICE_ENVS='{}'
    fi 

    if [ "$("${ZARUBA_HOME}/zaruba" isValidList "${_SERVICE_PORTS}")" -eq 0 ]
    then
        echo "ports ${_SERVICE_PORTS} is not a valid list, apply default value"
        _SERVICE_PORTS='[]'
    fi

    if [ "$("${ZARUBA_HOME}/zaruba" isValidList "${_DEPENDENCIES}")" -eq 0 ]
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

    _DEFAULT_PORT_LIST="$("${ZARUBA_HOME}/zaruba" getPortConfig "${_SERVICE_LOCATION}")"
    _DEFAULT_PORT_CONFIG="$("${ZARUBA_HOME}/zaruba" join "${_DEFAULT_PORT_LIST}" "\n")"

    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "${_REPLACEMENT_MAP}" \
        "zarubaImageName" "${_IMAGE_NAME}" \
        "zarubaContainerName" "${_CONTAINER_NAME}" \
        "zarubaServiceName" "${_SERVICE_NAME}" \
        "ZarubaServiceName" "${_PASCAL_SERVICE_NAME}" \
        "zaruba-service-name" "${_KEBAB_SERVICE_NAME}" \
        "ZARUBA_SERVICE_NAME" "${_UPPER_SNAKE_SERVICE_NAME}" \
        "zarubaStartCommand" "${_SERVICE_START_COMMAND}" \
        "zarubaServiceLocation" "$("${ZARUBA_HOME}/zaruba" getRelativePath "$(dirname "${_TASK_FILE_NAME}")" "${_SERVICE_LOCATION}")" \
        "zarubaRunnerVersion" "${_SERVICE_RUNNER_VERSION}" \
        "zarubaDefaultPortConfig" "${_DEFAULT_PORT_CONFIG}" \
    ) 

    "${ZARUBA_HOME}/zaruba" generate "${_TEMPLATE_LOCATION}" "${_DESTINATION}" "${_REPLACEMENT_MAP}"

    register_task_file "${_TASK_FILE_NAME}" "${_SERVICE_NAME}"

    "${ZARUBA_HOME}/zaruba" task addDependency ./main.zaruba.yaml "run${_PASCAL_SERVICE_NAME}" "${_DEPENDENCIES}"
    "${ZARUBA_HOME}/zaruba" setTaskEnv ./main.zaruba.yaml "run${_PASCAL_SERVICE_NAME}" "${_SERVICE_ENVS}"

            
    if [ "$("${ZARUBA_HOME}/zaruba" getListLength "${_SERVICE_PORTS}")" -gt 0 ]
    then
        _PORT_CONFIG_VALUE="$("${ZARUBA_HOME}/zaruba" join "${_SERVICE_PORTS}" )"
        _PORT_CONFIG="$("${ZARUBA_HOME}/zaruba" setMapElement "{}" "ports" "$_PORT_CONFIG_VALUE" )"
        "${ZARUBA_HOME}/zaruba" setTaskConfig ./main.zaruba.yaml "run${_PASCAL_SERVICE_NAME}" "${_PORT_CONFIG}"
    fi

}
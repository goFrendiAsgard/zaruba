. ${ZARUBA_HOME}/scripts/bash/util.sh
. ${ZARUBA_HOME}/scripts/bash/register_task_file.sh

# USAGE generate_service_task <template-location> <service-location> <service-name> <image-name> <container-name> <service-start-command> <service-runner-version> <service-ports> <service-envs> <dependencies>
generate_service_task() {
    TEMPLATE_LOCATION="${1}"
    SERVICE_LOCATION="${2}"
    SERVICE_NAME="${3}"
    IMAGE_NAME="${4}"
    CONTAINER_NAME="${5}"
    SERVICE_START_COMMAND="${6}"
    SERVICE_RUNNER_VERSION="${7}"
    SERVICE_PORTS="${8}"
    SERVICE_ENVS="${9}"
    DEPENDENCIES="${10}"

    DEFAULT_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" getServiceName "${SERVICE_LOCATION}")"
    SERVICE_NAME="$(get_value_or_default "${SERVICE_NAME}" "${DEFAULT_SERVICE_NAME}")"

    DEFAULT_IMAGE_NAME="$("${ZARUBA_HOME}/zaruba" strToKebab "${SERVICE_NAME}")"
    IMAGE_NAME="$(get_value_or_default "${IMAGE_NAME}" "${DEFAULT_IMAGE_NAME}")"

    DEFAULT_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" strToCamel "${SERVICE_NAME}")"
    CONTAINER_NAME="$(get_value_or_default "${CONTAINER_NAME}" "${DEFAULT_CONTAINER_NAME}")"

    PASCAL_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToPascal "${SERVICE_NAME}")"
    KEBAB_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToKebab "${SERVICE_NAME}")"
    SNAKE_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToSnake "${SERVICE_NAME}")"
    UPPER_SNAKE_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToUpper "${SNAKE_SERVICE_NAME}")"

    if [ "$({{ .Zaruba}} isValidMap "$SERVICE_ENVS")" -eq 0 ]
    then
        echo "${SERVICE_ENVS} is not a valid map"
        exit 1
    fi 

    if [ "$({{ .Zaruba}} isValidList "$SERVICE_PORTS")" -eq 0 ]
    then
        echo "${SERVICE_PORTS} is not a valid port"
        exit 1
    fi

    if [ "$({{ .Zaruba}} isValidList "$DEPENDENCIES")" -eq 0 ]
    then
        echo "${SERVICE_PORTS} is not a valid port"
        exit 1
    fi

    DESTINATION="./zaruba-task"
    TASK_FILE_NAME="${DESTINATION}/${SERVICE_NAME}.zaruba.yaml"
    if [ -f "${TASK_FILE_NAME}" ]
    then
        echo "file already exist: ${TASK_FILE_NAME}"
        exit 1
    fi

    DEFAULT_PORT_LIST="$("${ZARUBA_HOME}/zaruba" getPortConfig "${SERVICE_LOCATION}")"
    DEFAULT_PORT_CONFIG="$("${ZARUBA_HOME}/zaruba" join "${DEFAULT_PORT_LIST}" "\n")"

    REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zarubaImageName" "${IMAGE_NAME}" \
        "zarubaContainerName" "${CONTAINER_NAME}" \
        "zarubaServiceName" "${SERVICE_NAME}" \
        "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
        "zaruba-service-name" "${KEBAB_SERVICE_NAME}" \
        "ZARUBA_SERVICE_NAME" "${UPPER_SNAKE_SERVICE_NAME}" \
        "zarubaStartCommand" "${SERVICE_START_COMMAND}" \
        "zarubaServiceLocation" "$("${ZARUBA_HOME}/zaruba" getRelativePath "${DESTINATION}" "${SERVICE_LOCATION}")" \
        "zarubaRunnerVersion" "${SERVICE_RUNNER_VERSION}" \
        "zarubaDefaultPortConfig" "${DEFAULT_PORT_CONFIG}" \
    ) 

    "${ZARUBA_HOME}/zaruba" generate "${TEMPLATE_LOCATION}" "${DESTINATION}" "${REPLACEMENT_MAP}"

    . ${ZARUBA_HOME}/scripts/bash/register_task_file.sh
    register_task_file "${TASK_FILE_NAME}" "${SERVICE_NAME}"

    "${ZARUBA_HOME}/zaruba" addTaskDependency ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${DEPENDENCIES}"
    "${ZARUBA_HOME}/zaruba" setTaskEnv ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${SERVICE_ENVS}"

            
    if [ "$("${ZARUBA_HOME}/zaruba" getListLength "${SERVICE_PORTS}")" -gt 0 ]
    then
        PORT_CONFIG_VALUE="$("${ZARUBA_HOME}/zaruba" join "${SERVICE_PORTS}" )"
        PORT_CONFIG="$("${ZARUBA_HOME}/zaruba" setMapElement "{}" "ports" "$PORT_CONFIG_VALUE" )"
        "${ZARUBA_HOME}/zaruba" setTaskConfig ./main.zaruba.yaml "run${PASCAL_SERVICE_NAME}" "${PORT_CONFIG}"
    fi

}
. ${ZARUBA_HOME}/scripts/bash/util.sh
. ${ZARUBA_HOME}/scripts/bash/register_task_file.sh

# USAGE generate_docker_task <template-location> <image-name> <container-name> <service-name> <service-ports> <service-envs> <dependencies>
generate_docker_task() {
    TEMPLATE_LOCATION="${1}"
    IMAGE_NAME="${2}"
    CONTAINER_NAME="${3}"
    SERVICE_NAME="${4}"
    SERVICE_PORTS="${5}"
    SERVICE_ENVS="${6}"
    DEPENDENCIES="${7}"

    DEFAULT_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" strToCamel "${IMAGE_NAME}")"
    CONTAINER_NAME="$(get_value_or_default "${CONTAINER_NAME}" "${DEFAULT_CONTAINER_NAME}")"

    SERVICE_NAME="$(get_value_or_default "${SERVICE_NAME}" "${CONTAINER_NAME}")"

    PASCAL_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToPascal "${SERVICE_NAME}")"
    KEBAB_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" strToKebab "${SERVICE_NAME}")"

    if [ "$({{ .Zaruba}} isValidMap "$SERVICE_ENVS")" -eq 0 ]
    then
        echo "{{ $d.Red }}{{ $d.Bold }}${SERVICE_ENVS} is not a valid map{{ $d.Normal }}"
        exit 1
    fi 

    if [ "$({{ .Zaruba}} isValidList "$SERVICE_PORTS")" -eq 0 ]
    then
        echo "{{ $d.Red }}{{ $d.Bold }}${SERVICE_PORTS} is not a valid port{{ $d.Normal }}"
        exit 1
    fi

    if [ "$({{ .Zaruba}} isValidList "$DEPENDENCIES")" -eq 0 ]
    then
        echo "{{ $d.Red }}{{ $d.Bold }}${SERVICE_PORTS} is not a valid port{{ $d.Normal }}"
        exit 1
    fi

    DESTINATION="./zaruba-task"
    TASK_FILE_NAME="${DESTINATION}/${SERVICE_NAME}.zaruba.yaml"
    if [ -f "${TASK_FILE_NAME}" ]
    then
        echo "{{ $d.Red }}{{ $d.Bold }}file already exist: ${TASK_FILE_NAME}{{ $d.Normal }}"
        exit 1
    fi

    REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zarubaImageName" "${IMAGE_NAME}" \
        "zarubaContainerName" "${CONTAINER_NAME}" \
        "zarubaServiceName" "${SERVICE_NAME}" \
        "ZarubaServiceName" "${PASCAL_SERVICE_NAME}" \
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
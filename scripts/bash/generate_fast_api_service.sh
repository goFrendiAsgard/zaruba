. ${ZARUBA_HOME}/scripts/bash/util.sh
. ${ZARUBA_HOME}/scripts/bash/generate_service_task.sh

# USAGE generate_fast_api_service <service-template-location> <task-template-location> <service-name>
generate_fast_api_service() {
    _SERVICE_TEMPLATE_LOCATION="${1}"
    _TASK_TEMPLATE_LOCATION="${2}"
    _SERVICE_NAME="${3}"

    if [ -d "./${_SERVICE_NAME}" ]
    then
        echo "${_SERVICE_NAME} already exist"
        return
    fi

    _PASCAL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToPascal "${_SERVICE_NAME}")
    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" strToCamel "${_SERVICE_NAME}")
    _REPLACMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zarubaServiceName" "${_CAMEL_SERVICE_NAME}" \
        "ZarubaServiceName" "${_PASCAL_SERVICE_NAME}" \
    )

    echo "Creating Fast API Service: ${_SERVICE_NAME}"
    "${ZARUBA_HOME}/zaruba" generate "${_SERVICE_TEMPLATE_LOCATION}" . "${_REPLACMENT_MAP}"
    chmod 755 "${_CAMEL_SERVICE_NAME}/start.sh"

    if [ ! -f "./main.zaruba.yaml" ]
    then
        echo "Not in a project, skip creating shared-lib and task"
        return
    fi

    if [ ! -d "./shared-libs/python/helpers" ]
    then
        echo "Creating shared-lib directory"
        mkdir -p "./shared-libs/python/helpers"
        cp -rnT "./${_SERVICE_NAME}/helpers" "./shared-libs/python/helpers"
    fi

    echo "Creating shared-lib link for ${_SERVICE_NAME}"
    "${ZARUBA_HOME}/zaruba" setProjectValue "./default.values.yaml" "link::${_SERVICE_NAME}/helpers" "shared-libs/python/helpers"
    link_resource "shared-libs/python/helpers" "${_SERVICE_NAME}/helpers"

    _TASK_EXIST="$("${ZARUBA_HOME}/zaruba" isTaskExist "./main.zaruba.yaml" "run${_PASCAL_SERVICE_NAME}")"
    if [ "${_TASK_EXIST}" -eq 1 ]
    then
        echo "Service task already exist: run${_PASCAL_SERVICE_NAME}"
        return
    fi

    echo "Creating service task: ${_PASCAL_SERVICE_NAME}"
    _TASK_SERVICE_LOCATION="${_SERVICE_NAME}"
    _TASK_SERVICE_NAME="${_SERVICE_NAME}"
    _TASK_IMAGE_NAME=""
    _TASK_CONTAINER_NAME=""
    _TASK_SERVICE_START_COMMAND=""
    _TASK_SERVICE_RUNNER_VERSION=""
    _TASK_SERVICE_PORTS="[]"
    _TASK_SERVICE_ENVS="{}"
    _TASK_DEPENDENCIES="[]"
    _TASK_REPLACEMENT_MAP="{}"
    generate_service_task \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_TASK_SERVICE_LOCATION}" \
        "${_TASK_SERVICE_NAME}" \
        "${_TASK_IMAGE_NAME}" \
        "${_TASK_CONTAINER_NAME}" \
        "${_TASK_SERVICE_START_COMMAND}" \
        "${_TASK_SERVICE_RUNNER_VERSION}" \
        "${_TASK_SERVICE_PORTS}" \
        "${_TASK_SERVICE_ENVS}" \
        "${_TASK_DEPENDENCIES}" \
        "${_TASK_REPLACEMENT_MAP}"

}

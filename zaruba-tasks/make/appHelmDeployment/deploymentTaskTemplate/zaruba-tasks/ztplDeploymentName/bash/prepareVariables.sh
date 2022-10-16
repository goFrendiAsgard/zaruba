CONTAINER_ENVS='[]'
__OLD_IFS="${IFS}"
IFS=$'\n'
for KEY in $("${ZARUBA_BIN}" map rangeKey "${RAW_ENVS}")
do
    VAL="$("${ZARUBA_BIN}" map get "${RAW_ENVS}" "${KEY}")"
    DOUBLE_QUOTED_VAL="$("${ZARUBA_BIN}" str doubleQuote "${VAL}")"
    ENV_MAP='{}'
    ENV_MAP="$("${ZARUBA_BIN}" map set "${ENV_MAP}" "name" "${KEY}")"
    ENV_MAP="$("${ZARUBA_BIN}" map set "${ENV_MAP}" "value" "${DOUBLE_QUOTED_VAL}")"
    CONTAINER_ENVS="$("${ZARUBA_BIN}" list append "${CONTAINER_ENVS}" "${ENV_MAP}")"
done
IFS="${__OLD_IFS}"

CONTAINER_PORTS='[]'
SERVICE_PORTS='[]'
PORT_LIST=$("${ZARUBA_BIN}" str split "${ZARUBA_CONFIG_PORTS}")
for INDEX in $("${ZARUBA_BIN}" list rangeIndex "${PORT_LIST}")
do
    PORT_STR="$("${ZARUBA_BIN}" list get "${PORT_LIST}" "${INDEX}")"
    if [ -z "${PORT_STR}" ]
    then
        continue
    fi
    PORT_STR_PARTS="$("${ZARUBA_BIN}" str split "${PORT_STR}" ":")"
    echo "PARTS: ${PORT_STR_PARTS}"
    PORT_STR_PARTS_LENGTH=$("${ZARUBA_BIN}" list length "${PORT_STR_PARTS}")
    PORT="$("${ZARUBA_BIN}" list get "${PORT_STR_PARTS}" "$(( ${PORT_STR_PARTS_LENGTH} - 1 ))")"

    # add to service ports
    SERVICE_PORT_MAP='{"protocol": "TCP"}'
    SERVICE_PORT_MAP="$("${ZARUBA_BIN}" map set "${SERVICE_PORT_MAP}" "name" "port${INDEX}")"
    SERVICE_PORT_MAP="$("${ZARUBA_BIN}" map set "${SERVICE_PORT_MAP}" "targetPort" "port${INDEX}")"
    SERVICE_PORT_MAP="$("${ZARUBA_BIN}" map set "${SERVICE_PORT_MAP}" "port" "${PORT}")"
    SERVICE_PORTS="$("${ZARUBA_BIN}" list append "${SERVICE_PORTS}" "${SERVICE_PORT_MAP}")"

    # add to container ports
    CONTAINER_PORT_MAP='{"protocol": "TCP"}'
    CONTAINER_PORT_MAP="$("${ZARUBA_BIN}" map set "${CONTAINER_PORT_MAP}" "name" "port${INDEX}")"
    CONTAINER_PORT_MAP="$("${ZARUBA_BIN}" map set "${CONTAINER_PORT_MAP}" "containerPort" "${PORT}")"
    CONTAINER_PORTS="$("${ZARUBA_BIN}" list append "${CONTAINER_PORTS}" "${CONTAINER_PORT_MAP}")"
done

setDeploymentConfig "image.repository" "${IMAGE_REPOSITORY}"
setDeploymentConfig "image.tag" "${IMAGE_TAG}"
setDeploymentConfig "env" "${CONTAINER_ENVS}"
setDeploymentConfig "ports" "${CONTAINER_PORTS}"
setDeploymentConfig "service.ports" "${SERVICE_PORTS}"
_ZRB_CONTAINER_ENVS='[]'
for _ZRB_KEY in $("${ZARUBA_HOME}/zaruba" map rangeKey "${_ZRB_ENVS}")
do
    _ZRB_VAL="$("${ZARUBA_HOME}/zaruba" map get "${_ZRB_ENVS}" "${_ZRB_KEY}")"
    _ZRB_DOUBLE_QUOTED_VAL="$("${ZARUBA_HOME}/zaruba" str doubleQuote "${_ZRB_VAL}")"
    _ZRB_ENV_MAP='{}'
    _ZRB_ENV_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_ENV_MAP}" "name" "${_ZRB_KEY}")"
    _ZRB_ENV_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_ENV_MAP}" "value" "${_ZRB_DOUBLE_QUOTED_VAL}")"
    _ZRB_CONTAINER_ENVS="$("${ZARUBA_HOME}/zaruba" list append "${_ZRB_CONTAINER_ENVS}" "${_ZRB_ENV_MAP}")"
done

_ZRB_CONTAINER_PORTS='[]'
_ZRB_SERVICE_PORTS='[]'
_ZRB_PORT_LIST=$("${ZARUBA_HOME}/zaruba" str split "${_ZRB_RAW_CONFIG_PORTS}")
for _ZRB_INDEX in $("${ZARUBA_HOME}/zaruba" list rangeIndex "${_ZRB_PORT_LIST}")
do
    _ZRB_PORT_STR="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_PORT_LIST}" "${_ZRB_INDEX}")"
    if [ -z "${_ZRB_PORT_STR}" ]
    then
        continue
    fi
    _ZRB_PORT_STR_PARTS="$("${ZARUBA_HOME}/zaruba" str split "${_ZRB_PORT_STR}" ":")"
    echo "PARTS: ${_ZRB_PORT_STR_PARTS}"
    _ZRB_PORT_STR_PARTS_LENGTH=$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_PORT_STR_PARTS}")
    _ZRB_PORT="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_PORT_STR_PARTS}" "$(( ${_ZRB_PORT_STR_PARTS_LENGTH} - 1 ))")"

    # add to service ports
    _ZRB_SERVICE_PORT_MAP='{"protocol": "TCP"}'
    _ZRB_SERVICE_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_SERVICE_PORT_MAP}" "name" "port${_ZRB_INDEX}")"
    _ZRB_SERVICE_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_SERVICE_PORT_MAP}" "targetPort" "port${_ZRB_INDEX}")"
    _ZRB_SERVICE_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_SERVICE_PORT_MAP}" "port" "${_ZRB_PORT}")"
    _ZRB_SERVICE_PORTS="$("${ZARUBA_HOME}/zaruba" list append "${_ZRB_SERVICE_PORTS}" "${_ZRB_SERVICE_PORT_MAP}")"

    # add to container ports
    _ZRB_CONTAINER_PORT_MAP='{"protocol": "TCP"}'
    _ZRB_CONTAINER_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_CONTAINER_PORT_MAP}" "name" "port${_ZRB_INDEX}")"
    _ZRB_CONTAINER_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_CONTAINER_PORT_MAP}" "containerPort" "${_ZRB_PORT}")"
    _ZRB_CONTAINER_PORTS="$("${ZARUBA_HOME}/zaruba" list append "${_ZRB_CONTAINER_PORTS}" "${_ZRB_CONTAINER_PORT_MAP}")"
done

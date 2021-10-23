_ZRB_DEPLOYMENT_ENVS='[]'
for _ZRB_KEY in $("${ZARUBA_HOME}/zaruba" map rangeKey "${_ZRB_ENVS}")
do
    _ZRB_VAL="$("${ZARUBA_HOME}/zaruba" map get "${_ZRB_ENVS}" "${_ZRB_KEY}")"
    _ZRB_ENV_MAP='{}'
    _ZRB_ENV_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_ENV_MAP}" "name" "${_ZRB_KEY}")"
    _ZRB_ENV_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_ENV_MAP}" "value" "${_ZRB_VAL}")"
    _ZRB_DEPLOYMENT_ENVS="$("${ZARUBA_HOME}/zaruba" list append "${_ZRB_DEPLOYMENT_ENVS}" "${_ZRB_ENV_MAP}")"
done

_ZRB_CONTAINER_PORTS='[]'
_ZRB_PORT_LIST=$("${ZARUBA_HOME}/zaruba" str split "${_ZRB_RAW_CONFIG_PORTS}")
for _ZRB_INDEX in $("${ZARUBA_HOME}/zaruba" list rangeIndex "${_ZRB_PORT_LIST}")
do
    _ZRB_PORT_STR="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_PORT_LIST}" "${_ZRB_INDEX}")"
    _ZRB_PORT_STR_PARTS="$("${ZARUBA_HOME}/zaruba" str split "${_ZRB_PORT_STR}")"
    if [ "$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_PORT_STR_PARTS}")" = 2 ]
    then
        _ZRB_PORT="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_PORT_STR_PARTS}" 1)"
    else
        _ZRB_PORT="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_PORT_STR_PARTS}" 0)"
    fi

    _ZRB_CONTAINER_PORT_MAP='{"protocol": "TCP"}'
    _ZRB_CONTAINER_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_CONTAINER_PORT_MAP}" "name" "port${_ZRB_INDEX}")"
    _ZRB_CONTAINER_PORT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_CONTAINER_PORT_MAP}" "containerPort" "${_ZRB_PORT}")"
    _ZRB_CONTAINER_PORTS="$("${ZARUBA_HOME}/zaruba" list append "${_ZRB_CONTAINER_PORTS}" "${_ZRB_CONTAINER_PORT_MAP}")"
done

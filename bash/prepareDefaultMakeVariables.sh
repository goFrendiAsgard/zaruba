if [ -z "${_ZRB_APP_NAME}" ]
then
    _ZRB_APP_NAME="$("${ZARUBA_HOME}/zaruba" path getAppName "${_ZRB_APP_DIRECTORY}")"
fi

if [ -z "${_ZRB_APP_ENV_PREFIX}" ]
then
    _ZRB_APP_ENV_PREFIX_LOWER="$("${ZARUBA_HOME}/zaruba" str toSnake "${_ZRB_APP_NAME}")"
    _ZRB_APP_ENV_PREFIX="$("${ZARUBA_HOME}/zaruba" str toUpper "${_ZRB_APP_ENV_PREFIX_LOWER}")"
fi

if [ -z "${_ZRB_APP_IMAGE_NAME}" ]
then
    _ZRB_APP_IMAGE_NAME="$("${ZARUBA_HOME}/zaruba" str toKebab "${_ZRB_APP_DIRECTORY}")"
fi

if [ -z "${_ZRB_APP_CONTAINER_NAME}" ]
then
    _ZRB_APP_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" str toCamel "${_ZRB_APP_NAME}")"
fi

if [ -d "${_ZRB_APP_DIRECTORY}" ]
then
    echo "${_ZRB_APP_DIRECTORY}"
    _ZRB_DEFAULT_APP_ENVS="$("${ZARUBA_HOME}/zaruba" path getEnv "${_ZRB_APP_DIRECTORY}")"
    _ZRB_APP_ENVS="$("${ZARUBA_HOME}/zaruba" map merge "${_ZRB_APP_ENVS}" "${_ZRB_DEFAULT_APP_ENVS}")"

    _ZRB_DEFAULT_APP_PORTS="$("${ZARUBA_HOME}/zaruba" path getPortConfig "${_ZRB_APP_DIRECTORY}")"
    if [ "$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_APP_PORTS}")" = 0 ]
    then
        _ZRB_APP_PORTS="${_ZRB_DEFAULT_APP_PORTS}"
    fi
fi

_ZRB_APP_YAML_PORTS="$("${ZARUBA_HOME}/zaruba" list join "${_ZRB_APP_PORTS}")"

_ZRB_APP_MAP_ENVS='{}'
for _ZRB_KEY in $("${ZARUBA_HOME}/zaruba" map rangeKey "${_ZRB_APP_ENVS}")
do
    _ZRB_FROM="${_ZRB_APP_ENV_PREFIX}_${_ZRB_KEY}"
    _ZRB_DEFAULT="$("${ZARUBA_HOME}/zaruba" map get "${_ZRB_APP_ENVS}" "${_ZRB_KEY}")"
    _ZRB_APP_SINGLE_MAP_ENV='{}'
    _ZRB_APP_SINGLE_MAP_ENV="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_APP_SINGLE_MAP_ENV}" "from" "${_ZRB_FROM}")"
    _ZRB_APP_SINGLE_MAP_ENV="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_APP_SINGLE_MAP_ENV}" "default" "${_ZRB_DEFAULT}")"
    _ZRB_APP_MAP_ENVS="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_APP_MAP_ENVS}" "${_ZRB_KEY}" "${_ZRB_APP_SINGLE_MAP_ENV}")"
done
_ZRB_APP_YAML_ENVS="$("${ZARUBA_HOME}/zaruba" yaml print "${_ZRB_APP_MAP_ENVS}")"
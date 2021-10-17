# kafka envs
_ZRB_APP_KAFKA_ENVS="$("${ZARUBA_HOME}/zaruba" path getEnv "${_ZRB_APP_DIRECTORY}/kafka")"
# hosts values
_ZRB_APP_KAFKA_ENVS="$("${ZARUBA_HOME}/zaruba" map set "${_ZRB_APP_KAFKA_ENVS}" "KAFKA_CFG_ZOOKEEPER_CONNECT" "${_ZRB_APP_CONTAINER_NAME}Zookeeper")"
# yaml
_ZRB_APP_YAML_KAFKA_ENVS="$(_getYamlEnvs "${_ZRB_APP_KAFKA_ENVS}" "${_ZRB_APP_ENV_PREFIX}")"

# zookeeper envs
_ZRB_APP_ZOOKEEPER_ENVS="$("${ZARUBA_HOME}/zaruba" path getEnv "${_ZRB_APP_DIRECTORY}/zookeeper")"
# yaml
_ZRB_APP_YAML_ZOOKEEPER_ENVS="$(_getYamlEnvs "${_ZRB_APP_ZOOKEEPER_ENVS}" "${_ZRB_APP_ENV_PREFIX}_ZOOKEEPER")"

_ZRB_APP_KAFKA_YAML_PORTS="$("${ZARUBA_HOME}/zaruba" list join "${_ZRB_APP_KAFKA_PORTS}")"
_ZRB_APP_ZOOKEEPER_YAML_PORTS="$("${ZARUBA_HOME}/zaruba" list join "${_ZRB_APP_ZOOKEEPER_PORTS}")"
# kafka envs
_ZRB_APP_KAFKA_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/kafka")"
# yaml
_ZRB_APP_YAML_KAFKA_ENVS="$(_getYamlEnvs "${_ZRB_APP_KAFKA_ENVS}" "${_ZRB_APP_ENV_PREFIX}")"

# kafka connect envs
_ZRB_APP_CONNECT_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/kafkaConnect")"
# yaml
_ZRB_APP_YAML_CONNECT_ENVS="$(_getYamlEnvs "${_ZRB_APP_CONNECT_ENVS}" "${_ZRB_APP_ENV_PREFIX}_CONNECT")"

# kafka schema registry envs
_ZRB_APP_SCHEMA_REGISTRY_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/kafkaSchemaRegistry")"
# yaml
_ZRB_APP_YAML_SCHEMA_REGISTRY_ENVS="$(_getYamlEnvs "${_ZRB_APP_CONNECT_ENVS}" "${_ZRB_APP_ENV_PREFIX}_SCHEMA_REGISTRY")"

# akhq envs
_ZRB_APP_AKHQ_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/akhq")"
# yaml
_ZRB_APP_YAML_AKHQ_ENVS="$(_getYamlEnvs "${_ZRB_APP_AKHQ_ENVS}" "${_ZRB_APP_ENV_PREFIX}_AKHQ")"

# zookeeper envs
_ZRB_APP_ZOOKEEPER_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/zookeeper")"
# yaml
_ZRB_APP_YAML_ZOOKEEPER_ENVS="$(_getYamlEnvs "${_ZRB_APP_ZOOKEEPER_ENVS}" "${_ZRB_APP_ENV_PREFIX}_ZOOKEEPER")"

# ports
_ZRB_APP_YAML_AKHQ_PORTS="$("${ZARUBA_BIN}" list join "${_ZRB_APP_AKHQ_PORTS}")"
_ZRB_APP_YAML_KAFKA_PORTS="$("${ZARUBA_BIN}" list join "${_ZRB_APP_KAFKA_PORTS}")"
_ZRB_APP_YAML_CONNECT_PORTS="$("${ZARUBA_BIN}" list join "${_ZRB_APP_CONNECT_PORTS}")"
_ZRB_APP_YAML_SCHEMA_REGISTRY_PORTS="$("${ZARUBA_BIN}" list join "${_ZRB_APP_SCHEMA_REGISTRY_PORTS}")"
_ZRB_APP_YAML_ZOOKEEPER_PORTS="$("${ZARUBA_BIN}" list join "${_ZRB_APP_ZOOKEEPER_PORTS}")"
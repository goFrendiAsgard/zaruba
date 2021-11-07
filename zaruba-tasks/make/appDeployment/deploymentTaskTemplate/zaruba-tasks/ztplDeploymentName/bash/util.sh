setDeploymentConfig() {
    __ZRB_KEY="${1}"
    __ZRB_VAL="${2}"
    DEPLOYMENT_CONFIG="$("${ZARUBA_HOME}/zaruba" map set "${DEPLOYMENT_CONFIG}" "${__ZRB_KEY}" "${__ZRB_VAL}")"
}
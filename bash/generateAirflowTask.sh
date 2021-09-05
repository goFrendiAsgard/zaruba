. "${ZARUBA_HOME}/bash/generateDockerTask.sh"
. "${ZARUBA_HOME}/bash/generateServiceTask.sh"
. "${ZARUBA_HOME}/bash/setProjectValueUnlessExist.sh"

# USAGE generateServiceTask <template-location> <service-location> <service-name> <image-name> <container-name> <service-start-command> <service-runner-version> <service-ports> <service-envs> <dependencies> <replacement-map> <register-runner> <redis-template-location> <redis-service-name> <postgre-template-location> <postgre-service-name>
generateAirflowTask() {
    _AIRFLOW_TEMPLATE_LOCATION="${1}"
    _AIRFLOW_SERVICE_LOCATION="${2}"
    _AIRFLOW_SERVICE_NAME="${3}"
    _AIRFLOW_IMAGE_NAME="${4}"
    _AIRFLOW_CONTAINER_NAME="${5}"
    _AIRFLOW_START_COMMAND="${6}"
    _AIRFLOW_SERVICE_RUNNER_VERSION="${7}"
    _AIRFLOW_SERVICE_PORTS="${8}"
    _AIRFLOW_SERVICE_ENVS="${9}"
    _AIRFLOW_DEPENDENCIES="${10}"
    _AIRFLOW_REPLACEMENT_MAP="${11}"
    _AIRFLOW_REGISTER_RUNNER="${12}"
    _AIRFLOW_REDIS_TEMPLATE_LOCATION="${13}"
    _AIRFLOW_REDIS_SERVICE_NAME="${14}"
    _AIRFLOW_POSTGRE_TEMPLATE_LOCATION="${15}"
    _AIRFLOW_POSTGRE_SERVICE_NAME="${16}"

    # get redisTask and postgreTask
    _REDIS_TASK="run$("${ZARUBA_HOME}/zaruba" str toPascal "${_AIRFLOW_REDIS_SERVICE_NAME}")"
    _POSTGRE_TASK="run$("${ZARUBA_HOME}/zaruba" str toPascal "${_AIRFLOW_POSTGRE_SERVICE_NAME}")"

    # add to replacementMap
    _AIRFLOW_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_AIRFLOW_REPLACEMENT_MAP}" "zarubaRedisTask" "${_REDIS_TASK}" )"
    _AIRFLOW_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_AIRFLOW_REPLACEMENT_MAP}" "zarubaRedisService" "${_AIRFLOW_REDIS_SERVICE_NAME}" )"
    _AIRFLOW_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_AIRFLOW_REPLACEMENT_MAP}" "zarubaPostgreTask" "${_POSTGRE_TASK}" )"
    _AIRFLOW_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_AIRFLOW_REPLACEMENT_MAP}" "zarubaPostgreService" "${_AIRFLOW_POSTGRE_SERVICE_NAME}" )"


    if [ "$("${ZARUBA_HOME}/zaruba" task isExist ./main.zaruba.yaml "${_REDIS_TASK}")" = 0 ]
    then
        echo "create redis task: ${_REDIS_TASK}"
        generateDockerTask \
            "${_AIRFLOW_REDIS_TEMPLATE_LOCATION}" "" "${_AIRFLOW_REDIS_SERVICE_NAME}" \
            "" "[]" "{}" "[]" "{}" "1"
        setProjectValueUnlessExist redisServiceName "${_AIRFLOW_REDIS_AIRFLOW_SERVICE_NAME}"
    fi


    if [ "$("${ZARUBA_HOME}/zaruba" task isExist ./main.zaruba.yaml "${_POSTGRE_TASK}")" = 0 ]
    then
        echo "create postgre task: ${_POSTGRE_TASK}"
        generateDockerTask \
            "${_AIRFLOW_POSTGRE_TEMPLATE_LOCATION}" "" "${_AIRFLOW_POSTGRE_SERVICE_NAME}" \
            "" "[]" "{}" "[]" "{}" "1"
        setProjectValueUnlessExist postgreServiceName "${_AIRFLOW_POSTGRE_AIRFLOW_SERVICE_NAME}"
    fi

    generateServiceTask \
        "${_AIRFLOW_TEMPLATE_LOCATION}" \
        "${_AIRFLOW_SERVICE_LOCATION}" \
        "${_AIRFLOW_SERVICE_NAME}" \
        "${_AIRFLOW_IMAGE_NAME}" \
        "${_AIRFLOW_CONTAINER_NAME}" \
        "${_AIRFLOW_START_COMMAND}" \
        "${_AIRFLOW_SERVICE_RUNNER_VERSION}" \
        "${_AIRFLOW_SERVICE_PORTS}" \
        "${_AIRFLOW_SERVICE_ENVS}" \
        "${_AIRFLOW_DEPENDENCIES}" \
        "${_AIRFLOW_REPLACEMENT_MAP}" \
        "${_AIRFLOW_REGISTER_RUNNER}"
    setProjectValueUnlessExist airflowServiceName "${_AIRFLOW_SERVICE_NAME}"
}

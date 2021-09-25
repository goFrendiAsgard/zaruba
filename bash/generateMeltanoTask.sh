. "${ZARUBA_HOME}/bash/generateServiceTask.sh"
. "${ZARUBA_HOME}/bash/generateAirflowTask.sh"

# USAGE generateServiceTask <template-location> <service-location> <service-name> <image-name> <container-name> <service-start-command> <service-runner-version> <service-ports> <service-envs> <dependencies> <replacement-map> <register-runner> <airflow-template-location> <airflow-service-name> <redis-template-location> <redis-service-name> <postgre-template-location> <postgre-service-name>
generateMeltanoTask() {
    _MELTANO_TEMPLATE_LOCATION="${1}"
    _MELTANO_SERVICE_LOCATION="${2}"
    _MELTANO_SERVICE_NAME="${3}"
    _MELTANO_IMAGE_NAME="${4}"
    _MELTANO_CONTAINER_NAME="${5}"
    _MELTANO_START_COMMAND="${6}"
    _MELTANO_SERVICE_RUNNER_VERSION="${7}"
    _MELTANO_SERVICE_PORTS="${8}"
    _MELTANO_SERVICE_ENVS="${9}"
    _MELTANO_DEPENDENCIES="${10}"
    _MELTANO_REPLACEMENT_MAP="${11}"
    _MELTANO_REGISTER_RUNNER="${12}"
    _MELTANO_AIRFLOW_TEMPLATE_LOCATION="${13}"
    _MELTANO_AIRFLOW_SERVICE_NAME="${14}"
    _MELTANO_REDIS_TEMPLATE_LOCATION="${15}"
    _MELTANO_REDIS_SERVICE_NAME="${16}"
    _MELTANO_POSTGRE_TEMPLATE_LOCATION="${17}"
    _MELTANO_POSTGRE_SERVICE_NAME="${18}"

    # get airflowTask
    _MELTANO_AIRFLOW_PASCAL_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" str toPascal "${_MELTANO_AIRFLOW_SERVICE_NAME}")"
    _MELTANO_AIRFLOW_TASK="run${_MELTANO_AIRFLOW_PASCAL_SERVICE_NAME}"

    # add to replacementMap
    _MELTANO_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_MELTANO_REPLACEMENT_MAP}" "zarubaAirflowTask" "${_MELTANO_AIRFLOW_TASK}" )"
    _MELTANO_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_MELTANO_REPLACEMENT_MAP}" "zarubaAirflowService" "${_MELTANO_AIRFLOW_SERVICE_NAME}" )"
    _MELTANO_REPLACEMENT_MAP="$("${ZARUBA_HOME}/zaruba" map set "${_MELTANO_REPLACEMENT_MAP}" "ZarubaAirflowService" "${_MELTANO_AIRFLOW_PASCAL_SERVICE_NAME}" )"

    # create airflow task if needed
    if [ "$("${ZARUBA_HOME}/zaruba" task isExist ./main.zaruba.yaml "${_MELTANO_AIRFLOW_TASK}")" = 0 ]
    then
        echo "create airflow task: ${_MELTANO_AIRFLOW_TASK}"
        generateAirflowTask \
            "${_MELTANO_AIRFLOW_TEMPLATE_LOCATION}" \
            "${_MELTANO_AIRFLOW_SERVICE_NAME}" \
            "${_MELTANO_AIRFLOW_SERVICE_NAME}" \
            "" \
            "" \
            "" \
            "" \
            "[]" \
            "{}" \
            "[]" \
            "{}" \
            "1" \
            "${_MELTANO_REDIS_TEMPLATE_LOCATION}" \
            "${_MELTANO_REDIS_SERVICE_NAME}" \
            "${_MELTANO_POSTGRE_TEMPLATE_LOCATION}" \
            "${_MELTANO_POSTGRE_SERVICE_NAME}"
    fi

    # create meltano task
    generateServiceTask \
        "${_MELTANO_TEMPLATE_LOCATION}" \
        "${_MELTANO_SERVICE_LOCATION}" \
        "${_MELTANO_SERVICE_NAME}" \
        "${_MELTANO_IMAGE_NAME}" \
        "${_MELTANO_CONTAINER_NAME}" \
        "${_MELTANO_START_COMMAND}" \
        "${_MELTANO_SERVICE_RUNNER_VERSION}" \
        "${_MELTANO_SERVICE_PORTS}" \
        "${_MELTANO_SERVICE_ENVS}" \
        "${_MELTANO_DEPENDENCIES}" \
        "${_MELTANO_REPLACEMENT_MAP}" \
        "${_MELTANO_REGISTER_RUNNER}"
    
    # install meltano
    pip install meltano

    # create meltano project
    cd "${_MELTANO_SERVICE_LOCATION}"
    pipenv run meltano init app

    # add additional files to the project
    "${ZARUBA_HOME}/zaruba" generate "${ZARUBA_HOME}/templates/meltanoApp" ./app '{}'
    chmod 755 app/start.sh
}

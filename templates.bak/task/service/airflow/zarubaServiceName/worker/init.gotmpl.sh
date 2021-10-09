echo "Init zarubaServiceName worker"

setAirflowVar() {
    _KEY="${1}"
    _VAL="${2}"
    _AIRFLOW_VERSION="$(airflow version | cut -d'.' -f 1)"
    if [ "${_AIRFLOW_VERSION}" -eq 1 ]
    then
        airflow variables --set "${_KEY}" "${_VAL}"
    else
        airflow variables set "${_KEY}" "${_VAL}"
    fi
}
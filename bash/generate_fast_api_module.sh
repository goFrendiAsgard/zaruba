. ${ZARUBA_HOME}/bash/util.sh
. ${ZARUBA_HOME}/bash/generate_fast_api_service.sh

generate_fast_api_module() {
    _MODULE_TEMPLATE_LOCATION="${1}"
    _SERVICE_TEMPLATE_LOCATION="${2}"
    _TASK_TEMPLATE_LOCATION="${3}"
    _SERVICE_NAME="${4}"
    _MODULE_NAME="${5}"

    generate_fast_api_service \
        "${_SERVICE_TEMPLATE_LOCATION}" \
        "${_TASK_TEMPLATE_LOCATION}" \
        "${_SERVICE_NAME}"
    

    if [ -d "./${_SERVICE_NAME}/${_MODULE_NAME}" ]
    then
        echo "${_SERVICE_NAME}/${_MODULE_NAME} already exist"
        return
    fi


    echo "Creating Fast API module: ${_SERVICE_NAME}/${_MODULE_NAME}"
    _PASCAL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str pascal "${_SERVICE_NAME}")
    _CAMEL_SERVICE_NAME=$("${ZARUBA_HOME}/zaruba" str camel "${_SERVICE_NAME}")
    _PASCAL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str pascal "${_MODULE_NAME}")
    _CAMEL_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str camel "${_MODULE_NAME}")
    _SNAKE_MODULE_NAME=$("${ZARUBA_HOME}/zaruba" str snake "${_MODULE_NAME}")
    _REPLACEMENT_MAP=$("${ZARUBA_HOME}/zaruba" setMapElement "{}" \
        "zarubaServiceName" "${_CAMEL_SERVICE_NAME}" \
        "ZarubaServiceName" "${_PASCAL_SERVICE_NAME}" \
        "zarubaModuleName" "${_CAMEL_MODULE_NAME}" \
        "ZarubaModuleName" "${_PASCAL_MODULE_NAME}" \
        "zaruba_module_name" "${_SNAKE_MODULE_NAME}" \
    )
    "${ZARUBA_HOME}/zaruba" generate "${_MODULE_TEMPLATE_LOCATION}/zarubaServiceName" "${_CAMEL_SERVICE_NAME}" "${_REPLACEMENT_MAP}"

    # get main.py lines
    _MAIN_LINES=$("${ZARUBA_HOME}/zaruba" readLines "${_CAMEL_SERVICE_NAME}/main.py")

    # import module
    _IMPORT_MODULE_PARTIAL=$(cat "${_MODULE_TEMPLATE_LOCATION}/partials/import_module.py")
    _IMPORT_MODULE_PARTIAL=$("${ZARUBA_HOME}/zaruba" str replace "${_IMPORT_MODULE_PARTIAL}" "${_REPLACEMENT_MAP}")
    _IMPORT_MODULE_LINES=$("${ZARUBA_HOME}/zaruba" str split "${_IMPORT_MODULE_PARTIAL}")

    # load module
    _LOAD_MODULE_PARTIAL=$(cat "${_MODULE_TEMPLATE_LOCATION}/partials/load_module.py")
    _LOAD_MODULE_PARTIAL=$("${ZARUBA_HOME}/zaruba" str replace "${_LOAD_MODULE_PARTIAL}" "${_REPLACEMENT_MAP}")
    _LOAD_MODULE_LINES=$("${ZARUBA_HOME}/zaruba" str split "${_LOAD_MODULE_PARTIAL}")

    # update main.py
    _MAIN_LINES=$("${ZARUBA_HOME}/zaruba" list merge "${_IMPORT_MODULE_LINES}" "${_MAIN_LINES}" "${_LOAD_MODULE_LINES}")
    "${ZARUBA_HOME}/zaruba" writeLines "${_CAMEL_SERVICE_NAME}/main.py" "${_MAIN_LINES}"

}

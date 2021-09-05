# USAGE getDockerContainerName <container-name> <image-name> <template-location>
getDockerContainerName() {
    _CONTAINER_NAME="${1}"
    _IMAGE_NAME="${2}"
    _TEMPLATE_LOCATION="${3}"
    if [ -z "${_CONTAINER_NAME}" ]
    then
        _DEFAULT_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" str toCamel "${_IMAGE_NAME}")"
        if [ -z "${_DEFAULT_CONTAINER_NAME}" ]
        then
            _DEFAULT_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" path getServiceName "${_TEMPLATE_LOCATION}")"
        fi
        echo "${_DEFAULT_CONTAINER_NAME}"
    else
        echo "${_CONTAINER_NAME}"
    fi
}

# USAGE getDockerServiceName <service-name> <container-name>
getDockerServiceName() {
    _SERVICE_NAME="${1}"
    _CONTAINER_NAME="${2}"
    if [ -z "${_SERVICE_NAME}" ]
    then
        echo "${_CONTAINER_NAME}"
    else
        echo "${_SERVICE_NAME}"
    fi
}

# USAGE getServiceName <service-name> <service-location>
getServiceName() {
    _SERVICE_NAME="${1}"
    _SERVICE_LOCATION="${2}"
    if [ -z "${_SERVICE_NAME}" ]
    then
        _DEFAULT_SERVICE_NAME="$("${ZARUBA_HOME}/zaruba" path getServiceName "${_SERVICE_LOCATION}")"
        echo "${_DEFAULT_SERVICE_NAME}"
    else
        echo "${_SERVICE_NAME}"
    fi
}

# USAGE getServiceImageName <image-name> <service-name>
getServiceImageName() {
    _IMAGE_NAME="${1}"
    _SERVICE_NAME"${2}"
    if [ -z "${_IMAGE_NAME}" ]
    then
        _DEFAULT_IMAGE_NAME="$("${ZARUBA_HOME}/zaruba" str toKebab "${_SERVICE_NAME}")"
        echo "${_DEFAULT_IMAGE_NAME}"
    else
        echo "${_IMAGE_NAME}"
    fi
}

# USAGE getContainerName <container-name> <service-name>
getServiceContainerName() {
    _CONTAINER_NAME="${1}"
    _SERVICE_NAME="${2}"
    if [ -z "${CONTAINER_NAME}" ]
    then
        _DEFAULT_CONTAINER_NAME="$("${ZARUBA_HOME}/zaruba" str toCamel "${_SERVICE_NAME}")"
        echo "${_DEFAULT_CONTAINER_NAME}"
    else
        echo "${_CONTAINER_NAME}"
    fi
}


#!/bin/sh
set -e
PROJECT_DIR=$1
SERVICE_NAME=$2
TEMPLATE_DIR=$(pwd)
EXISTING_SERVICE_COUNT=$(find ${PROJECT_DIR}/services -maxdepth 1 -type d|wc -l)
HTTP_PORT=$(expr ${EXISTING_SERVICE_COUNT} + 3010)
HTTP_TEST_PORT=$(expr ${EXISTING_SERVICE_COUNT} + 4010)

# get service name
if [ -z ${SERVICE_NAME} ]
then
    read -p "Type your service name: " SERVICE_NAME
fi
SERVICE_NAME=$(basename ${SERVICE_NAME})
UPPER_SERVICE_NAME=$(echo ${SERVICE_NAME} | tr '[a-z]' '[A-Z]')

# if service already exists, we should not override it
if [ -e ${PROJECT_DIR}/services/${SERVICE_NAME} ]
then
    echo "Service ${SERVICE} already exists in ${PROJECT_DIR}/services"
    exit 0
fi

SERVICE_DIR=${PROJECT_DIR}/services/${SERVICE_NAME}

# copyAndMoveToService()
#   copy "service" folder from TEMPLATE_DIR to SERVICE_DIR
copyAndMoveToService() {
    cp -r ${TEMPLATE_DIR}/service ${SERVICE_DIR}
    cd ${SERVICE_DIR}
}

# substituteValue(_FILES...)
#   substitute SERVICENAME, servicename, 3010, and 4010 in all given _FILES
substituteValues() {
    for _FILE in ${@}
    do
        sed -i "s/SERVICENAME/${UPPER_SERVICE_NAME}/g" ${SERVICE_DIR}/${_FILE}
        sed -i "s/servicename/${SERVICE_NAME}/g" ${SERVICE_DIR}/${_FILE}
        sed -i "s/3010/${HTTP_PORT}/g" ${SERVICE_DIR}/${_FILE}
        sed -i "s/4010/${HTTP_TEST_PORT}/g" ${SERVICE_DIR}/${_FILE}
    done
}

# copySharedLibs(_SHARED_LIB_DIR, _SRC_DIR, _HELPER_LIST...)
#   copy to shared libs in case of the shared lib doesn't exists
copySharedLibs() {
    _SHARED_LIB_DIR=${1}
    _SRC_DIR=${2}
    shift
    shift
    _HELPER_LIST=${@}
    mkdir -p ${_SHARED_LIB_DIR}
    for _HELPER in ${_HELPER_LIST}
    do
        if [ -e ${_SHARED_LIB_DIR}/${_HELPER} ]
        then
            echo "${_SHARED_LIB_DIR}/${_HELPER} is already exists"
        else
            cp -r ${SERVICE_DIR}/${_SRC_DIR}/${_HELPER} ${_SHARED_LIB_DIR}/${_HELPER}
        fi
    done
}

# adjustZarubaYaml()
#   adjust zaruba yaml name
adjustZarubaYaml() {
    mv ${SERVICE_DIR}/zaruba.yaml ${SERVICE_DIR}/${SERVICE_NAME}.zaruba.yaml 
}

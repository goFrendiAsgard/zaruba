#!/bin/sh
set -e

if [ -z ${ZARUBA_TEST_DIR} ]
then
    echo "[ERROR] ZARUBA_TEST_DIR is not defined"
    exit 1
fi

if [ -z ${ZARUBA_TEMPLATE_DIR} ]
then
    echo "[ERROR] ZARUBA_TEMPLATE_DIR is not defined"
    exit 1
fi

go build
# tear down
rm -Rf ${ZARUBA_TEST_DIR}
rm -Rf ${ZARUBA_TEMPLATE_DIR}
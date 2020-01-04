#!/bin/sh
set -x

go build
# tear down
rm -Rf ${ZARUBA_TEST_DIR}
rm -Rf ${ZARUBA_TEMPLATE_DIR}
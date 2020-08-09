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

# check things
git version
docker version
go version

# check compilation
go build

# clean up directories
rm -Rf ${ZARUBA_TEST_DIR}
rm -Rf ${ZARUBA_TEMPLATE_DIR}

# prepare directories
mkdir -p ${ZARUBA_TEST_DIR}
cp -R ./templates ${ZARUBA_TEMPLATE_DIR}

# prepare server-repos
mkdir -p test-resource/server-repos/parent
cd test-resource/server-repos/parent
rm -Rf .git && rm -Rf *
git init
cd ..
mkdir -p child
cd child
rm -Rf .git && rm -Rf *
git init
cd ../..
#!/bin/sh
set -e

# check whether user provide `--y` parameter or not
YES_TO_ALL=0
for PARAMETER in "$@" ; do
    if [ ${PARAMETER} = "--y" ] ||[ ${PARAMETER} = "--Y" ]
    then
        YES_TO_ALL=1
        break
    fi
done

# Set default value for ZARUBA_TEMPLATE_DIR
if [ -z ${ZARUBA_TEMPLATE_DIR} ]
then
    ZARUBA_TEMPLATE_DIR=${HOME}/.zaruba/templates
fi

echo "* CHECK GIT VERSION"
git version

echo "* CHECK DOCKER VERSION"
docker version

echo "* CHECK GO VERSION"
go version

echo "* INSTALLING ZARUBA"
go get -u github.com/state-alchemists/zaruba

echo "* CREATING ZARUBA CONFIG DIRECTORY"
if [ -z ${HOME} ]
then
    HOME=~
fi
mkdir -p ${HOME}/.zaruba
mkdir -p ${HOME}/.zaruba/templates

# create zaruba/zaruba.env
echo 'if [ -z ${HOME} ]' > ${HOME}/.zaruba/zaruba.env
echo 'then' >> ${HOME}/.zaruba/zaruba.env
echo '    HOME=~' >> ${HOME}/.zaruba/zaruba.env
echo 'fi' >> ${HOME}/.zaruba/zaruba.env
echo 'export PATH="$(go env GOPATH)/bin:${PATH}"' >> ${HOME}/.zaruba/zaruba.env
echo 'export ZARUBA_SHELL="/bin/bash"' >> ${HOME}/.zaruba/zaruba.env
echo 'export ZARUBA_SHELL_ARG="-c"' >> ${HOME}/.zaruba/zaruba.env
echo 'export ZARUBA_TEMPLATE_DIR="${HOME}/.zaruba/templates"' >> ${HOME}/.zaruba/zaruba.env

# create hook for bash and zsh
for CONFIG_FILE in .bashrc .zshrc
do
    if [ -e ${HOME}/${CONFIG_FILE} ]
    then
        if $(grep -q ".zaruba/zaruba.env" ${HOME}/${CONFIG_FILE})
        then
            echo "* ZARUBA HOOK FOR ${CONFIG_FILE} ALREADY EXISTS"
        else
            echo "* ADD ZARUBA HOOK FOR ${CONFIG_FILE}"
            echo '' >> ${HOME}/${CONFIG_FILE}
            echo '# init zaruba' >> ${HOME}/${CONFIG_FILE}
            echo 'source ${HOME}/.zaruba/zaruba.env' >> ${HOME}/${CONFIG_FILE}
            echo '' >> ${HOME}/${CONFIG_FILE}
        fi
    fi
done


if [ $YES_TO_ALL -eq 0 ]
then
    read -p "Do you want to install pre-built templates(Y/n)? " INSTALL_TEMPLATE
fi
if [ $YES_TO_ALL -eq 1 ] || [ $INSTALL_TEMPLATE = "Y" ] || [ $INSTALL_TEMPLATE = "y" ]
then
    TEMPLATE_PATH=$(go env GOPATH)/src/github.com/state-alchemists/zaruba/templates
    if [ -e templates ]
    then
        TEMPLATE_PATH=templates
    fi
    echo "* INSTALL PRE-BUILT TEMPLATES"
    for COMPONENT in $(ls ${TEMPLATE_PATH})
    do
        echo "   - ${COMPONENT}"
        [ -e ${HOME}/.zaruba/templates/${COMPONENT} ] && rm -Rf ${HOME}/.zaruba/templates/${COMPONENT}
        cp -R ${TEMPLATE_PATH}/${COMPONENT} ${HOME}/.zaruba/templates
        chmod 755 -R ${HOME}/.zaruba/templates/${COMPONENT}
    done
fi

echo "* DONE"
echo "Please invoke 'source ${HOME}/.zaruba/zaruba.env' or restart this terminal in order to start using zaruba"

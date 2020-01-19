#!/bin/sh

echo "* INSTALLING ZARUBA"
go get -u github.com/state-alchemists/zaruba

echo "* CREATING ZARUBA CONFIG DIRECTORY"
mkdir -p ${HOME}/zaruba
mkdir -p ${HOME}/zaruba/template

# create zaruba/zaruba.env
echo 'export PATH="$(go env GOPATH)/bin:${PATH}"' > ${HOME}/zaruba/zaruba.env
echo 'export ZARUBA_SHELL="/bin/bash"' >> ${HOME}/zaruba/zaruba.env
echo 'export ZARUBA_SHELL_ARG="-c"' >> ${HOME}/zaruba/zaruba.env
echo 'export ZARUBA_TEMPLATE_DIR="${HOME}/zaruba/template"' >> ${HOME}/zaruba/zaruba.env

# create hooks
echo "* ADD ZARUBA HOOK FOR bash"
echo '# init zaruba' >> ${HOME}/.bashrc
echo 'source ${HOME}/zaruba/zaruba.env' >> ${HOME}/.bashrc
echo "* ADD ZARUBA HOOK FOR zsh"
echo '# init zaruba' >> ${HOME}/.zshrc
echo 'source ${HOME}/zaruba/zaruba.env' >> ${HOME}/.zshrc

echo "* DONE"
echo "Please invoke 'source ${HOME}/zaruba/zaruba.env' or restart this terminal in order to start using zaruba"
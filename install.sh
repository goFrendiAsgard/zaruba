#!/bin/sh

set -e

echo 💀 Cloning Zaruba 
rm -Rf "${HOME}/.zaruba"
git clone --depth 1 https://github.com/state-alchemists/zaruba "${HOME}/.zaruba"

echo 💀 Build Zaruba 
echo "Build Zaruba"
cd "${HOME}/.zaruba"
go build

echo 💀 Remove old Zaruba 
go clean -i github.com/state-alchemists/zaruba

if echo "${PATH}" | grep "${HOME}/.zaruba"
then
    echo 💀 PATH is already containing '${HOME}/.zaruba'
else
    echo 💀 Injecting '${HOME}/.zaruba' to PATH
    PATH=$PATH:"${HOME}/.zaruba"
    if [ -e "${HOME}/.bashrc" ]
    then
        echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.bashrc"
    fi
    if [ -e "${HOME}/.zshrc" ]
    then
        echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.zshrc"
    fi
fi
echo 💀 Installation success 

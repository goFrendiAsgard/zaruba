#!/bin/sh

set +e
echo 💀 Remove old Zaruba 
go clean -i github.com/state-alchemists/zaruba
rm -Rf "${HOME}/.zaruba"

set -e
echo 💀 Cloning Zaruba 
git clone --depth 1 https://github.com/state-alchemists/zaruba "${HOME}/.zaruba"

echo 💀 Build Zaruba 
cd "${HOME}/.zaruba"
go build


if [ -f /usr/bin/zaruba ]
then
    echo 💀 Remove old '/usr/bin/zaruba' symlink
    sudo rm -Rf /usr/bin/zaruba
fi

set +e
echo 💀 Create '/usr/bin/zaruba' symlink
sudo ln -s ${HOME}/.zaruba/zaruba /usr/bin/zaruba

if [ "$?" = 0 ]
then
    echo 💀 '/usr/bin/zaruba' symlink created
else
    set -e
    echo 💀 Failed to create symlink, injecting PATH instead
    if echo "${PATH}" | grep "${HOME}/.zaruba"
    then
        echo 💀 PATH is already containing '${HOME}/.zaruba'
    else
        echo 💀 Injecting '${HOME}/.zaruba' to PATH
        PATH=$PATH:"${HOME}/.zaruba"
        if [ -e "${HOME}/.bashrc" ]
        then
            echo 💀 Injecting '${HOME}/.zaruba' to .bashrc
            echo "" >> "${HOME}/.bashrc"
            echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.bashrc"
        fi
        if [ -e "${HOME}/.zshrc" ]
        then
            echo 💀 Injecting '${HOME}/.zaruba' to .zshrc
            echo "" >> "${HOME}/.zshrc"
            echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.zshrc"
        fi
    fi
fi

echo 💀 Installation success 
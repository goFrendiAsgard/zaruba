set +e
echo "💀 Remove old Zaruba"
go clean -i github.com/state-alchemists/zaruba
if [ -f "${HOME}/.zaruba/scripts/bash/bootstrap.sh" ]
then
    echo "💀 Backup bootstrap script"
    cp "${HOME}/.zaruba/scripts/bash/bootstrap.sh" "${HOME}/zaruba-bootstrap.sh.bak"
fi
rm -Rf "${HOME}/.zaruba"

set -e
echo "💀 Cloning Zaruba"
git clone --depth 1 https://github.com/state-alchemists/zaruba "${HOME}/.zaruba"
if [ -f "${HOME}/zaruba-bootstrap.sh.bak" ]
then
    echo "💀 Restore bootstrap script"
    mv "${HOME}/zaruba-bootstrap.sh.bak" "${HOME}/.zaruba/scripts/bash/bootstrap.sh"
fi

echo "💀 Build Zaruba"
cd "${HOME}/.zaruba"
git fetch --tags
go build


if [ -f /usr/bin/zaruba ]
then
    echo "💀 Remove old '/usr/bin/zaruba' symlink"
    sudo rm -Rf /usr/bin/zaruba
fi

set +e
echo "💀 Create '/usr/bin/zaruba' symlink"
sudo ln -s ${HOME}/.zaruba/zaruba /usr/bin/zaruba

if [ "$?" = 0 ]
then
    set -e
    echo "💀 '/usr/bin/zaruba' symlink created"
else
    set -e
    echo "💀 Failed to create symlink, injecting PATH instead"
    if echo "${PATH}" | grep "${HOME}/.zaruba"
    then
        echo "💀 PATH is already containing '${HOME}/.zaruba'"
    else
        echo "💀 Injecting '${HOME}/.zaruba' to PATH"
        PATH=$PATH:"${HOME}/.zaruba"
        if [ -e "${HOME}/.bashrc" ]
        then
            echo "💀 Injecting '${HOME}/.zaruba' to .bashrc"
            echo "" >> "${HOME}/.bashrc"
            echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.bashrc"
        fi
        if [ -e "${HOME}/.zshrc" ]
        then
            echo "💀 Injecting '${HOME}/.zaruba' to .zshrc"
            echo "" >> "${HOME}/.zshrc"
            echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.zshrc"
        fi
    fi
fi

echo "💀 Installation success"

echo "💀 zaruba can install several packages in case of you are using ubuntu"
read -p "💀 Do you want to setup ubuntu? (y/N): "  CONFIRMATION
if [ "${CONFIRMATION}" = "y" ]
then
    zaruba please setupUbuntu -i
    echo "💀 Ubuntu setup success"
else
    echo "💀 You can perform this task later by invoking 'zaruba please setupUbuntu -i'"
fi

echo ""
echo "💀 pyenv lets you easily switch between multiple versions of Python"
echo "💀 Some of zaruba's tasks depend on Python"
read -p "💀 Do you want to setup pyenv? (y/N): "  CONFIRMATION
if [ "${CONFIRMATION}" = "y" ]
then
    zaruba please setupPyenv -i
    echo "💀 Pyenv setup success"
else
    echo "💀 You can perform this task later by invoking 'zaruba please setupPyenv -i'"
fi

echo ""
echo "💀 nvm is a version manager for node.js"
read -p "💀 Do you want to setup nvm? (y/N): "  CONFIRMATION
if [ "${CONFIRMATION}" = "y" ]
then
    zaruba please setupNvm -i
    echo 💀 Nvm setup success 
else
    echo 💀 You can perform this task later by invoking 'zaruba please setupNvm -i' 
fi

echo ""
echo "💀 When you setup pyenv/nvm, zaruba will also make bootstrap script on '~/.zaruba/scripts/bootstrap.sh' to be used internally"
read -p "💀 Do you want to also inject the bootstrap script into your bash/zsh? (y/N): "  CONFIRMATION
if [ "${CONFIRMATION}" = "y" ]
then
    zaruba please injectBootstrap -i
    echo 💀 Bootstrap injected 
else
    echo 💀 You can perform this task later by invoking 'zaruba please injectBootstrap' 
fi


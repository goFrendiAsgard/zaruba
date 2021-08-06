set +e
echo "💀 Removing old zaruba installation"
rm -Rf "${HOME}/.zaruba"

set -e
echo "💀 Cloning zaruba source code"
git clone --depth 1 https://github.com/state-alchemists/zaruba "${HOME}/.zaruba"

echo "💀 Building zaruba"
cd "${HOME}/.zaruba"
git fetch --tags
go build

echo "💀 Injecting zaruba to the PATH"
if echo "${PATH}" | grep '${HOME}/.zaruba'
then
    echo "💀 PATH is already containing '${HOME}/.zaruba'"
else
    echo "💀 Injecting '${HOME}/.zaruba' to PATH"
    if [ -f "${HOME}/.profile" ]
    then
        echo "💀 Injecting '${HOME}/.zaruba' to .profile"
        echo "" >> "${HOME}/.profile"
        echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.profile"
    fi
    PATH=$PATH:"${HOME}/.zaruba"
fi
echo "💀 Zaruba is"


echo "💀 Installation success"
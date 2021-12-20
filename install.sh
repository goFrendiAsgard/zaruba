set -e
GIT_URL="https://github.com/state-alchemists/zaruba"
INSTALLATION_DIR="${HOME}/.zaruba"
OLD_SYMLINK="/usr/bin/zaruba"
INIT_FILE="${INSTALLATION_DIR}/init.sh"
BACKUP_INIT_FILE="${HOME}/zaruba.init.sh.bak"
LOAD_INIT_FILE_SCRIPT='
if [ "${ZARUBA_INIT_SCRIPT_LOADED}" != "1" ] && [ -f "${HOME}/.zaruba/init.sh" ]
then
    . "${HOME}/.zaruba/init.sh"
fi
'
BRANCH="${1}"
if [ -z "${BRANCH}" ]
then
    BRANCH="master"
fi

# Remove symlink, since 0.7.0 zaruba doesn't need symlink
echo "💀 Removing old zaruba installation."
if [ -f "${OLD_SYMLINK}" ]
then
    echo "💀 Removing symlink."
    sudo rm "${OLD_SYMLINK}"
fi

# Backup init file
if [ -f "${INIT_FILE}" ]
then
    echo "💀 Backing up init.sh."
    cp "${INIT_FILE}" "${BACKUP_INIT_FILE}"
fi

# Remove old zaruba installation
if [ -d "${INSTALLATION_DIR}" ]
then
    echo "💀 Removing old installation folder."
    rm -Rf "${INSTALLATION_DIR}"
fi

# Clone from repo
echo "💀 Cloning zaruba source code."
git clone --depth 1 --branch "${BRANCH}" "${GIT_URL}" "${INSTALLATION_DIR}"
if [ "${BRANCH}" != "master" ]
then
    git checkout -b "master"
    git checkout "${BRANCH}"
fi

# Build
echo "💀 Building zaruba."
cd "${INSTALLATION_DIR}"
chmod 755 buildZaruba.sh
. ./buildZaruba.sh
chmod -R 755 "${INSTALLATION_DIR}/setup"

# Restore init script or create a new one
if [ -f "${BACKUP_INIT_FILE}" ]
then
    echo "💀 Restoring init.sh."
    mv "${BACKUP_INIT_FILE}" "${INIT_FILE}" 
else
    echo "💀 Creating init.sh."
    cp "${INSTALLATION_DIR}/setup/templates/bash/init.sh" "${INIT_FILE}"
    chmod 755 "${INIT_FILE}"
fi

# Inject init script to user's terminal
echo "💀 Injecting init script."
if echo "${PATH}" | grep -q '${HOME}/.zaruba'
then
    echo "💀 PATH is already containing '${HOME}/.zaruba'."
else
    for FILE in "${HOME}/.profile" "${HOME}/.bashrc" "${HOME}/.zshrc" 
    do
        if [ -f "${FILE}" ]
        then
            if grep -q ".zaruba/init.sh" "${FILE}"
            then
                echo "💀 ${FILE} is already containing load init file script."
            else
                echo "💀 Injecting init script to ${FILE}."
                echo "${LOAD_INIT_FILE_SCRIPT}" >> "${FILE}"
            fi
        fi
    done
fi

# Starting init script for current session
echo "💀 Loading init script."
. "${INSTALLATION_DIR}/init.sh"

echo "🎉🎉🎉"
echo "💀 Installation success."

# Run third party installer
echo "💀 You can now setup/install third party tools by invoking zaruba install <command>"
"${HOME}/.zaruba/zaruba install --help"
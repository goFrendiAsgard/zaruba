set -e

GIT_URL="https://github.com/state-alchemists/zaruba"
INSTALLATION_DIR="${HOME}/.zaruba"
OLD_SYMLINK="/usr/bin/zaruba"
INIT_FILE="${INSTALLATION_DIR}/init.sh"
BACKUP_INIT_FILE="${HOME}/zaruba.init.sh.bak"
LOAD_INIT_FILE_SCRIPT='
if [ -f "${HOME}/.zaruba/init.sh"]
then
    . "${HOME}/.zaruba/init.sh"
fi
'

# Remove symlink, since 0.7.0 zaruba doesn't need symlink
echo "💀 Removing old zaruba installation."
if [ -f "${OLD_SYMLINK}" ]
then
    echo "💀 Removing symlink."
    sudo rm "${OLD_SYMLINK}"
fi

# Backup init file
if [ -f "" ]
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
git clone --depth 1 "${GIT_URL}" "${INSTALLATION_DIR}"

# Build
echo "💀 Building zaruba."
cd "${INSTALLATION_DIR}"
git fetch --tags
go build

# Restore init script or create a new one
if [ -f "${BACKUP_INIT_FILE}" ]
then
    echo "💀 Restoring init.sh."
    mv "${BACKUP_INIT_FILE}" "${INIT_FILE}" 
else
    echo "💀 Creating init.sh."
    cp "${INSTALLATION_DIR}/templates/bash/init.sh" "${INIT_FILE}"
    chmod 755 "${INIT_FILE}"
fi

# Inject init script to user's terminal
echo "💀 Injecting init script."
if echo "${PATH}" | grep '${HOME}/.zaruba'
then
    echo "💀 PATH is already containing '${HOME}/.zaruba'."
else
    for FILE in "${HOME}/.profile" "${HOME}/.bashrc" "${HOME}/.zshrc" 
    then
        if [ -f "${FILE}" ]
        then
            echo "💀 Injecting init script to ${FILE}."
            echo "${LOAD_INIT_FILE_SCRIPT}" >> "${FILE}/"
        fi
    fi
fi
echo "💀 Starting init script."
. "${INSTALLATION_DIR}/init.sh"

echo "🎉🎉🎉"
echo "💀 Installation success."
echo "💀 You can now setup/install additional packages."

chmod 755 -R "${INSTALLATION_DIR}/setup"


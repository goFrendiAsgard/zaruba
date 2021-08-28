set -e
SETUP_PATH="$(dirname $(realpath "${0}"))"

echo "🔽 Downloading and installing nvm."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash

echo "🔽 Update init script."
SCRIPT="$(cat "${SETUP_PATH}/../templates/bash/nvm.sh")"
echo "${SCRIPT}" >> "${SETUP_PATH}/../init.sh"
. "${SETUP_PATH}/../init.sh"

echo "🎉🎉🎉"
echo "Nvm installed, make sure to update your shell configuration"